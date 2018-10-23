package p2p

import (
	"context"

	"github.com/aergoio/aergo/pkg/component"
	"github.com/aergoio/aergo/types"
	"github.com/golang/protobuf/proto"
	crypto "github.com/libp2p/go-libp2p-crypto"
	ifconnmgr "github.com/libp2p/go-libp2p-interface-connmgr"
	inet "github.com/libp2p/go-libp2p-net"
	peer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	protocol "github.com/libp2p/go-libp2p-protocol"
	ma "github.com/multiformats/go-multiaddr"
	msmux "github.com/multiformats/go-multistream"
	"github.com/stretchr/testify/mock"
)

type mockHost struct {
	psStub pstore.Peerstore
}

func (m mockHost) Peerstore() pstore.Peerstore {
	return m.psStub
}

func (m mockHost) ID() peer.ID {
	return "mockk"
}

func (m mockHost) Addrs() []ma.Multiaddr {
	return make([]ma.Multiaddr, 0, 0)
}

// Networks returns the Network interface of the Host
func (m mockHost) Network() inet.Network {
	return nil
}

// Mux returns the Mux multiplexing incoming streams to protocol handlers
func (m mockHost) Mux() *msmux.MultistreamMuxer {
	return nil
}

// Connect ensures there is a connection between this host and the peer with
// given peer.ID. Connect will absorb the addresses in pi into its internal
// pstore. If there is not an active connection, Connect will issue a
// h.Network.Dial, and block until a connection is open, or an error is
// returned. // TODO: Relay + NAT.
func (m mockHost) Connect(ctx context.Context, pi pstore.PeerInfo) error {
	return nil
}

// SetStreamHandler sets the protocol handler on the Host's Mux.
// This is equivalent to:
//   host.Mux().SetHandler(proto, handler)
// (Threadsafe)
func (m mockHost) SetStreamHandler(pid protocol.ID, handler inet.StreamHandler) {
	// DO nothing
}

// SetStreamHandlerMatch sets the protocol handler on the Host's Mux
// using a matching function for protocol selection.
func (m mockHost) SetStreamHandlerMatch(protocol.ID, func(string) bool, inet.StreamHandler) {

}

// RemoveStreamHandler removes a handler on the mux that was set by
// SetStreamHandler
func (m mockHost) RemoveStreamHandler(pid protocol.ID) {}

// NewStream opens a new stream to given peer p, and writes a p2p/protocol
// header with given protocol.ID. If there is no connection to p, attempts
// to create one. If ProtocolID is "", writes no header.
// (Threadsafe)
func (m mockHost) NewStream(ctx context.Context, p peer.ID, pids ...protocol.ID) (inet.Stream, error) {
	return nil, nil
}

// Close shuts down the host, its Network, and services.
func (m mockHost) Close() error {
	return nil
}

func (m mockHost) ConnManager() ifconnmgr.ConnManager {
	return nil
}

// MockPeerManager is an autogenerated mock type for the MockPeerManager type
type MockPeerManager struct {
	mock.Mock
}

// AddNewPeer provides a mock function with given fields: _a0
func (_m *MockPeerManager) AddNewPeer(_a0 PeerMeta) {
	_m.Called(_a0)
}

// RemovePeer provides a mock function with given fields: _a0
func (_m *MockPeerManager) RemovePeer(_a0 peer.ID) {
	_m.Called(_a0)
}

// NotifyPeerHandshake provides a mock function with given fields: _a0
func (_m *MockPeerManager) NotifyPeerHandshake(_a0 peer.ID) {
	_m.Called(_a0)
}

// NotifyPeerAddressReceived provides a mock function with given fields: _a0
func (_m *MockPeerManager) NotifyPeerAddressReceived(_a0 []PeerMeta) {
	_m.Called(_a0)
}

// Addrs provides a mock function with given fields:
func (_m *MockPeerManager) Addrs() []ma.Multiaddr {
	ret := _m.Called()

	var r0 []ma.Multiaddr
	if rf, ok := ret.Get(0).(func() []ma.Multiaddr); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ma.Multiaddr)
		}
	}

	return r0
}

// AuthenticateMessage provides a mock function with given fields: message, data
func (_m *MockPeerManager) AuthenticateMessage(message proto.Message, data *types.MsgHeader) bool {
	ret := _m.Called(message, data)

	var r0 bool
	if rf, ok := ret.Get(0).(func(proto.Message, *types.MsgHeader) bool); ok {
		r0 = rf(message, data)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *MockPeerManager) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConnManager provides a mock function with given fields:
func (_m *MockPeerManager) ConnManager() ifconnmgr.ConnManager {
	ret := _m.Called()

	var r0 ifconnmgr.ConnManager
	if rf, ok := ret.Get(0).(func() ifconnmgr.ConnManager); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ifconnmgr.ConnManager)
		}
	}

	return r0
}

// Connect provides a mock function with given fields: ctx, pi
func (_m *MockPeerManager) Connect(ctx context.Context, pi pstore.PeerInfo) error {
	ret := _m.Called(ctx, pi)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, pstore.PeerInfo) error); ok {
		r0 = rf(ctx, pi)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPeer provides a mock function with given fields: ID
func (_m *MockPeerManager) LookupPeer(ID peer.ID) (*remotePeerImpl, bool) {
	ret := _m.Called(ID)

	var r0 *remotePeerImpl
	if rf, ok := ret.Get(0).(func(peer.ID) *remotePeerImpl); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*remotePeerImpl)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(peer.ID) bool); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetPeer provides a mock function with given fields: ID
func (_m *MockPeerManager) GetPeer(ID peer.ID) (RemotePeer, bool) {
	ret := _m.Called(ID)

	var r0 RemotePeer
	if rf, ok := ret.Get(0).(func(peer.ID) RemotePeer); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(RemotePeer)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(peer.ID) bool); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetPeers provides a mock function with given fields:
func (_m *MockPeerManager) GetPeers() []RemotePeer {
	ret := _m.Called()

	var r0 []RemotePeer
	if rf, ok := ret.Get(0).(func() []RemotePeer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]RemotePeer)
		}
	}

	return r0
}

// GetPeerAddresses provides a mock function with given fields:
func (_m *MockPeerManager) GetPeerAddresses() ([]*types.PeerAddress, []types.PeerState) {
	ret := _m.Called()

	var r0 []*types.PeerAddress
	if rf, ok := ret.Get(0).(func() []*types.PeerAddress); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.PeerAddress)
		}
	}

	var r1 []types.PeerState
	if rf, ok := ret.Get(1).(func() []types.PeerState); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).([]types.PeerState)
	}

	return r0, r1
}

// GetStatus provides a mock function with given fields:
func (_m *MockPeerManager) GetStatus() component.Status {
	ret := _m.Called()

	var r0 component.Status
	if rf, ok := ret.Get(0).(func() component.Status); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(component.Status)
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *MockPeerManager) ID() peer.ID {
	ret := _m.Called()

	var r0 peer.ID
	if rf, ok := ret.Get(0).(func() peer.ID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(peer.ID)
	}

	return r0
}

// Mux provides a mock function with given fields:
func (_m *MockPeerManager) Mux() *msmux.MultistreamMuxer {
	ret := _m.Called()

	var r0 *msmux.MultistreamMuxer
	if rf, ok := ret.Get(0).(func() *msmux.MultistreamMuxer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msmux.MultistreamMuxer)
		}
	}

	return r0
}

// Network provides a mock function with given fields:
func (_m *MockPeerManager) Network() inet.Network {
	ret := _m.Called()

	var r0 inet.Network
	if rf, ok := ret.Get(0).(func() inet.Network); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(inet.Network)
		}
	}

	return r0
}

// NewMessageData provides a mock function with given fields: messageID, gossip
func (_m *MockPeerManager) NewMessageData(messageID string, gossip bool) *types.MsgHeader {
	ret := _m.Called(messageID, gossip)

	var r0 *types.MsgHeader
	if rf, ok := ret.Get(0).(func(string, bool) *types.MsgHeader); ok {
		r0 = rf(messageID, gossip)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.MsgHeader)
		}
	}

	return r0
}

// NewStream provides a mock function with given fields: ctx, p, pids
func (_m *MockPeerManager) NewStream(ctx context.Context, p peer.ID, pids ...protocol.ID) (inet.Stream, error) {
	_va := make([]interface{}, len(pids))
	for _i := range pids {
		_va[_i] = pids[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, p)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 inet.Stream
	if rf, ok := ret.Get(0).(func(context.Context, peer.ID, ...protocol.ID) inet.Stream); ok {
		r0 = rf(ctx, p, pids...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(inet.Stream)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, peer.ID, ...protocol.ID) error); ok {
		r1 = rf(ctx, p, pids...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Peerstore provides a mock function with given fields:
func (_m *MockPeerManager) Peerstore() pstore.Peerstore {
	ret := _m.Called()

	var r0 pstore.Peerstore
	if rf, ok := ret.Get(0).(func() pstore.Peerstore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pstore.Peerstore)
		}
	}

	return r0
}

// PrivateKey provides a mock function with given fields:
func (_m *MockPeerManager) PrivateKey() crypto.PrivKey {
	ret := _m.Called()

	var r0 crypto.PrivKey
	if rf, ok := ret.Get(0).(func() crypto.PrivKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.PrivKey)
		}
	}

	return r0
}

// PublicKey provides a mock function with given fields:
func (_m *MockPeerManager) PublicKey() crypto.PubKey {
	ret := _m.Called()

	var r0 crypto.PubKey
	if rf, ok := ret.Get(0).(func() crypto.PubKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.PubKey)
		}
	}

	return r0
}

// RemoveStreamHandler provides a mock function with given fields: pid
func (_m *MockPeerManager) RemoveStreamHandler(pid protocol.ID) {
	_m.Called(pid)
}

// SelfNodeID provides a mock function with given fields:
func (_m *MockPeerManager) SelfMeta() PeerMeta {
	ret := _m.Called()

	var r0 PeerMeta
	if rf, ok := ret.Get(0).(func() PeerMeta); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(PeerMeta)
	}

	return r0
}

// SelfNodeID provides a mock function with given fields:
func (_m *MockPeerManager) SelfNodeID() peer.ID {
	ret := _m.Called()

	var r0 peer.ID
	if rf, ok := ret.Get(0).(func() peer.ID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(peer.ID)
	}

	return r0
}

// SetStreamHandler provides a mock function with given fields: pid, handler
func (_m *MockPeerManager) SetStreamHandler(pid protocol.ID, handler inet.StreamHandler) {
	_m.Called(pid, handler)
}

// SetStreamHandlerMatch provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockPeerManager) SetStreamHandlerMatch(_a0 protocol.ID, _a1 func(string) bool, _a2 inet.StreamHandler) {
	_m.Called(_a0, _a1, _a2)
}

// SignProtoMessage provides a mock function with given fields: message
func (_m *MockPeerManager) SignProtoMessage(message proto.Message) ([]byte, error) {
	ret := _m.Called(message)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(proto.Message) []byte); ok {
		r0 = rf(message)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(proto.Message) error); ok {
		r1 = rf(message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields:
func (_m *MockPeerManager) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *MockPeerManager) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMessageData provides a mock function with given fields: messageID, gossip
func (_m *MockPeerManager) HandleNewBlockNotice(peerID peer.ID, hash BlkHash, data *types.NewBlockNotice) {
	_m.Called(peerID, hash, data)
}

// NewMessageData provides a mock function with given fields: messageID, gossip
func (_m *MockPeerManager) HandleNewTxNotice(peerID peer.ID, hash []TxHash, data *types.NewTransactionsNotice) {
	_m.Called(peerID, hash, data)
}
