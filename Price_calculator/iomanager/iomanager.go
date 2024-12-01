package iomanager

// Creating a interface so that the prices can accept either of cmd of filemanager
// as they both have same methods or rather they should have same method
type IoManager interface {
	ReadLines() ([]string, error)
	WriteJson(data interface{}) error
}
