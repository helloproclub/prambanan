package errors

var (
	// ErrGeneral something's when wrong
	ErrGeneral = NewError(404, 1000, "Terjadi kesalahan")
	// ErrDecode failed to decode request
	ErrDecodeRequest = NewError(422, 1000, "Parameter yang anda masukkan tidak bisa kami baca")
	// ErrNikInvalid NIK seems not to be an integer
	ErrNikInvalid = NewError(422, 1001, "NIK yang anda masukkan tidak bisa kami baca")
	// ErrGender gender code is invalid
	ErrGender = NewError(422, 1001, "Jenis kelamin dan tanggal lahir, gagal diketahui")
	// ErrMonth month code is invalid
	ErrMonth = NewError(422, 1002, "Tanggal lahir gagal diketahui.")
)
