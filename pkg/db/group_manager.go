package db

import "sync"

// GroupManager 管理数据库客户端Group
type GroupManager struct {
	mu     sync.RWMutex
	groups map[string]*Group
}

// Add 添加Group到Manager
func (gm *GroupManager) Add(name string, group *Group) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	gm.groups[name] = group
}

// Get 从Manage获取Group
func (gm *GroupManager) Get(name string) *Group {
	gm.mu.RLock()
	defer gm.mu.RUnlock()
	return gm.groups[name]
}

// NewGroupManager 实例化GroupManager
func NewGroupManager() *GroupManager {
	return &GroupManager{
		groups: make(map[string]*Group),
	}
}

// GroupManager实例
var groupManager = NewGroupManager()

// Get 从GroupManager获取Group
func Get(name string) *Group {
	return groupManager.Get(name)
}
