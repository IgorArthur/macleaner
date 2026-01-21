package tests

type MockFS struct {
	existingPaths []string
	removedPaths  []string
}

func (m *MockFS) ExpandPath(path string) (string, error) {
	return path, nil
}

func (m *MockFS) DirSize(path string) (int64, error) {
	return 0, nil
}

func (m *MockFS) RemoveAll(path string) error {
	m.removedPaths = append(m.removedPaths, path)
	return nil
}

func (m *MockFS) Exists(path string) bool {
	for _, p := range m.existingPaths {
		if p == path {
			return true
		}
	}
	return false
}
