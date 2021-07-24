package encryption_service_wrapper

type EncryptedResponseContract struct {
	Status bool          `json:"status"`
	Error  interface{}   `json:"error"`
	Result EncryptedData `json:"result"`
}
type EncryptedData struct {
	Data []byte `json:"data"`
}
