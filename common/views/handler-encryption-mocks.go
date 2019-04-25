package views

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
	"github.com/pydio/cells/common/crypto"
	"github.com/pydio/cells/common/proto/encryption"
	"github.com/pydio/cells/idm/key"
	"strings"
)

const (
	aesGCMTagSize   = 16
	blockHeaderSize = 12
	nodeKeySize     = 32
)

type mockSetNodeInfoStream struct {
	inStream  chan interface{}
	outStream chan interface{}

	keys   map[string]*encryption.NodeKey
	blocks map[string][]*encryption.Block

	cursor int

	exchangeError error
}

func newMockSendInfoStream(keys map[string]*encryption.NodeKey, blocks map[string][]*encryption.Block) *mockSetNodeInfoStream {
	return &mockSetNodeInfoStream{
		outStream: make(chan interface{}, 1),
		inStream:  make(chan interface{}, 1),
		keys:      keys,
		blocks:    blocks,
		cursor:    0,
	}
}

func (sc *mockSetNodeInfoStream) getClient() encryption.NodeKeyManager_SetNodeInfoClient {
	return &mockSendBlockStreamClient{
		outStream: sc.inStream,
		inStream:  sc.outStream,
	}
}

func (sc *mockSetNodeInfoStream) exchange() {
	for sc.exchangeError == nil {
		var req encryption.SetNodeInfoRequest
		sc.exchangeError = sc.RecvMsg(&req)
		if sc.exchangeError != nil {
			return
		}

		if req.NodeKey != nil {
			sc.keys[req.NodeId] = req.NodeKey
		}

		if req.Block != nil {
			nodeBlocks := sc.blocks[req.NodeId]
			if nodeBlocks == nil {
				nodeBlocks = []*encryption.Block{}
			}
			nodeBlocks = append(nodeBlocks, req.Block)
			sc.blocks[req.NodeId] = nodeBlocks
		}

		rsp := &encryption.SetNodeInfoResponse{}
		sc.exchangeError = sc.SendMsg(rsp)
	}
}

func (sc *mockSetNodeInfoStream) SendMsg(msg interface{}) error {
	sc.outStream <- msg
	return nil
}

func (sc *mockSetNodeInfoStream) RecvMsg(msgi interface{}) error {
	o := <-sc.inStream
	inMsg := o.(*encryption.SetNodeInfoRequest)
	msg := msgi.(*encryption.SetNodeInfoRequest)
	msg.NodeKey = inMsg.NodeKey
	msg.Block = inMsg.Block
	msg.NodeId = inMsg.NodeId
	return nil
}

func (sc *mockSetNodeInfoStream) Close() error {
	close(sc.inStream)
	close(sc.outStream)
	return nil
}

func (sc *mockSetNodeInfoStream) Send(request *encryption.SetNodeInfoRequest) error {
	return sc.SendMsg(request)
}

// MockSendBlockClient
type mockSendBlockStreamClient struct {
	inStream  chan interface{}
	outStream chan interface{}
}

func (sc *mockSendBlockStreamClient) SendMsg(msg interface{}) error {
	sc.outStream <- msg
	return nil
}

func (sc *mockSendBlockStreamClient) RecvMsg(msgi interface{}) error {
	_ = <-sc.inStream
	return nil
}

func (sc *mockSendBlockStreamClient) Close() error {
	close(sc.outStream)
	return nil
}

func (sc *mockSendBlockStreamClient) Send(request *encryption.SetNodeInfoRequest) error {
	return sc.SendMsg(request)
}

// NodeKeyManagerClient
type mockNodeKeyManagerClient struct {
	keys   map[string]*encryption.NodeKey
	blocks map[string][]*encryption.Block
}

func NewMockNodeKeyManagerClient() encryption.NodeKeyManagerClient {
	return &mockNodeKeyManagerClient{
		keys:   map[string]*encryption.NodeKey{},
		blocks: map[string][]*encryption.Block{},
	}
}

func (m *mockNodeKeyManagerClient) GetNodeInfo(ctx context.Context, in *encryption.GetNodeInfoRequest, opts ...client.CallOption) (*encryption.GetNodeInfoResponse, error) {

	nodeKey, entryFound := m.keys[in.NodeId]
	if !entryFound {
		return nil, errors.NotFound("mock.NodeKeyManager", "Key not found")
	}

	rsp := &encryption.GetNodeInfoResponse{
		NodeInfo: &encryption.NodeInfo{
			Node: &encryption.Node{
				NodeId: in.NodeId,
				Legacy: false,
			},
			NodeKey: nodeKey,
		},
	}

	if in.WithRange {
		foundEncryptedOffset := false
		foundEncryptedLimit := in.PlainLength > 0

		encryptedOffset := int64(0)
		encryptedLimit := int64(0)
		currentPlainOffset := int64(0)
		currentPlainLength := int64(0)

		blocks, foundBlocks := m.blocks[in.NodeId]
		if foundBlocks {
			for _, b := range blocks {

				plainBlockSize := b.BlockSize - aesGCMTagSize
				encryptedBlockSize := b.BlockSize + b.HeaderSize

				encryptedLimit += int64(encryptedBlockSize)

				if !foundEncryptedOffset {
					left := in.PlainOffset - currentPlainOffset
					if left == 0 {
						foundEncryptedOffset = true
						rsp.HeadSKippedPlainBytesCount = 0

					} else if left <= int64(plainBlockSize) {
						foundEncryptedOffset = true
						rsp.HeadSKippedPlainBytesCount = in.PlainOffset - currentPlainOffset
						currentPlainLength = int64(plainBlockSize) - rsp.HeadSKippedPlainBytesCount

						if currentPlainLength >= in.PlainLength {
							foundEncryptedLimit = true
							break
						}
						continue
					} else {
						currentPlainOffset += int64(plainBlockSize)
						encryptedOffset += int64(encryptedBlockSize)
					}
				}

				if foundEncryptedOffset && !foundEncryptedLimit {
					if currentPlainLength+int64(plainBlockSize) >= in.PlainLength {
						foundEncryptedLimit = true
					}
				}
			}
		}

		rsp.EncryptedOffset = encryptedOffset
		rsp.EncryptedCount = encryptedLimit - encryptedOffset
	}
	return rsp, nil
}

func (m *mockNodeKeyManagerClient) SetNodeInfo(ctx context.Context, opts ...client.CallOption) (encryption.NodeKeyManager_SetNodeInfoClient, error) {
	stream := newMockSendInfoStream(m.keys, m.blocks)
	go stream.exchange()
	return stream.getClient(), nil
}

func (m *mockNodeKeyManagerClient) DeleteNode(ctx context.Context, in *encryption.DeleteNodeRequest, opts ...client.CallOption) (*encryption.DeleteNodeResponse, error) {
	var entriesToDelete []string

	for entry, _ := range m.keys {
		if strings.HasSuffix(entry, fmt.Sprintf(":%s", in.NodeId)) {
			entriesToDelete = append(entriesToDelete, entry)
		}
	}
	for _, entry := range entriesToDelete {
		delete(m.keys, entry)
		delete(m.blocks, entry)
	}
	return &encryption.DeleteNodeResponse{}, nil
}

func (m *mockNodeKeyManagerClient) DeleteNodeKey(ctx context.Context, in *encryption.DeleteNodeKeyRequest, opts ...client.CallOption) (*encryption.DeleteNodeKeyResponse, error) {
	entry := fmt.Sprintf("%s:%s", in.UserId, in.NodeId)
	delete(m.keys, entry)
	delete(m.blocks, entry)
	return &encryption.DeleteNodeKeyResponse{}, nil
}

func (m *mockNodeKeyManagerClient) DeleteNodeSharedKey(ctx context.Context, in *encryption.DeleteNodeSharedKeyRequest, opts ...client.CallOption) (*encryption.DeleteNodeSharedKeyResponse, error) {
	entry := fmt.Sprintf("shared:%s:%s", in.UserId, in.NodeId)
	delete(m.keys, entry)
	delete(m.blocks, entry)
	return &encryption.DeleteNodeSharedKeyResponse{}, nil
}

// mockUserKeyTool
type mockUserKeyTool struct {
	key []byte
}

func NewMockUserKeyTool() key.UserKeyTool {
	keyByte := make([]byte, 16)
	_, _ = rand.Read(keyByte)
	return &mockUserKeyTool{
		key: keyByte,
	}
}

func (ukt *mockUserKeyTool) GetEncrypted(ctx context.Context, keyID string, data []byte) ([]byte, error) {
	return crypto.Seal(ukt.key, data)
}

func (ukt *mockUserKeyTool) GetDecrypted(ctx context.Context, keyID string, data []byte) ([]byte, error) {
	return crypto.Open(ukt.key, data[:12], data[12:])
}
