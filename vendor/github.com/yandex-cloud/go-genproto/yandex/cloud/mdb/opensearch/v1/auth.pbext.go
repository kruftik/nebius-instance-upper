// Code generated by protoc-gen-goext. DO NOT EDIT.

package opensearch

func (m *AuthSettings) SetSaml(v *SAMLSettings) {
	m.Saml = v
}

func (m *SAMLSettings) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *SAMLSettings) SetIdpEntityId(v string) {
	m.IdpEntityId = v
}

func (m *SAMLSettings) SetIdpMetadataFile(v []byte) {
	m.IdpMetadataFile = v
}

func (m *SAMLSettings) SetSpEntityId(v string) {
	m.SpEntityId = v
}

func (m *SAMLSettings) SetDashboardsUrl(v string) {
	m.DashboardsUrl = v
}

func (m *SAMLSettings) SetRolesKey(v string) {
	m.RolesKey = v
}

func (m *SAMLSettings) SetSubjectKey(v string) {
	m.SubjectKey = v
}
