package model

import "math/rand"

type User struct {
	Id                 int    `json:"id"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	Email              string `json:"email"`
	AlamatKTP          string `json:"alamat_ktp"`
	Pekerjaan          string `json:"pekerjaan"`
	NamaLengkap        string `json:"nama_lengkap"`
	PendidikanTerakhir string `json:"pendidikan_terakhir"`
	NomorTelepon       string `json:"nomor_telepon"`
}

func (u User) Validation(c string) map[string]string {
	var val = make(map[string]string)
	switch c {
	case "login":
		if u.Username == "" || u.Email == "" {
			val["username"] = "Username atau Email tidak boleh kosong"
		} else if u.Password == "" {
			val["password"] = "Password Tidak boleh kosong"
		}
	case "create":
		if u.Username == "" {
			val["username"] = "Username Tidak Boleh Kosong"
		} else if u.Password == "" {
			val["password"] = "Password Tidak Boleh Kosong"
		} else if u.Email == "" {
			val["email"] = "Email Tidak Boleh Kosong"
		} else if u.AlamatKTP == "" {
			val["alamat_ktp"] = "Alamat KTP Tidak Boleh Kosong"
		} else if u.Pekerjaan == "" {
			val["pekerjaan"] = "Pekerjaan Tidak Boleh Kosong"
		} else if u.NamaLengkap == "" {
			val["nama_lengkap"] = "Nama Lengkap Tidak Boleh Kosong"
		} else if u.PendidikanTerakhir == "" {
			val["pendidikan_terakhir"] = "Pendidikan Terakhir Tidak Boleh Kosong"
		} else if u.NomorTelepon == "" {
			val["nomor_telepon"] = "Nomor Telepon Tidak Boleh Kosong"
		}
	}
	return val
}

func InitUser() (User, User) {
	var user User
	user.Id = rand.Int()
	user.Username = "agam"
	user.Password = "agam"
	user.NamaLengkap = "Firman Agam"
	user.Email = "email@gmail.com"
	user.Pekerjaan = "idle"
	user.PendidikanTerakhir = "S1"
	user.NomorTelepon = "6281214025919"
	user.AlamatKTP = "jalan raya"
	var user1 User
	user1.Id = rand.Int()
	user1.Username = "firda"
	user1.Password = "firda"
	user1.NamaLengkap = "Firda"
	user1.Email = "email@gmail.com"
	user1.Pekerjaan = "idle"
	user1.PendidikanTerakhir = "S1"
	user1.NomorTelepon = "0129942312"
	user1.AlamatKTP = "pekayon"
	return user, user1
}
