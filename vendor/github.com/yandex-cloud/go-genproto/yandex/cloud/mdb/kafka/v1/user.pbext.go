// Code generated by protoc-gen-goext. DO NOT EDIT.

package kafka

func (m *User) SetName(v string) {
	m.Name = v
}

func (m *User) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *User) SetPermissions(v []*Permission) {
	m.Permissions = v
}

func (m *UserSpec) SetName(v string) {
	m.Name = v
}

func (m *UserSpec) SetPassword(v string) {
	m.Password = v
}

func (m *UserSpec) SetPermissions(v []*Permission) {
	m.Permissions = v
}

func (m *Permission) SetTopicName(v string) {
	m.TopicName = v
}

func (m *Permission) SetRole(v Permission_AccessRole) {
	m.Role = v
}

func (m *Permission) SetAllowHosts(v []string) {
	m.AllowHosts = v
}
