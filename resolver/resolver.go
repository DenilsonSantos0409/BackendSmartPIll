package resolver

import (
	"errors"
	db "smartpill/connection"
	model "smartpill/models"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Buat user baru
func CreateUser(username, email, password string) model.User {
	// Hash password sebelum disimpan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// Jika gagal hashing, kembalikan password asli (atau bisa juga panik/log error)
		hashedPassword = []byte(password)
	}

	user := model.User{
		Username:  username,
		Email:     email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}

	db.DB.Create(&user)
	return user
}

// Ambil semua user
func GetAllUser() []model.User {
	var user []model.User
	db.DB.Find(&user)
	return user
}

// Ambil satu user berdasarkan ID
func GetUserByID(id uint) (model.User, error) {
	var user model.User
	result := db.DB.First(&user, id)
	return user, result.Error
}

func GetAllObat() []model.Obat {
	var obat []model.Obat
	db.DB.Find(&obat)
	return obat
}

func GetObatByUser(userId uint) []model.Obat {
	var obat []model.Obat
	db.DB.Where("user_id = ?", userId).Find(&obat)
	return obat
}

func UpdateUser(id uint, username, email, password, profil *string) (*model.User, error) {
	var user model.User

	// Cari data user berdasarkan ID
	if err := db.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	// Update hanya jika nilainya tidak nil (biar bisa optional)
	if username != nil {
		user.Username = *username
	}
	if email != nil {
		user.Email = *email
	}
	if password != nil {
		// Hash password baru
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}
	if profil != nil {
		user.Profil = profil
	}

	user.CreatedAt = time.Now() // opsional: update waktu

	if err := db.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateObat(userId uint, nama, dosis, frekuensi, catatan, tanggal, waktuStr string) model.Obat {
	obat := model.Obat{
		User_id:   int(userId),
		Nama_Obat: nama,
		Dosis:     dosis,
		Frekuensi: frekuensi,
		Catatan:   catatan,
		Tanggal:   tanggal,
		Waktu:     waktuStr,
		CreatedAt: time.Now(),
	}
	db.DB.Create(&obat)
	return obat
}

func UpdateObat(id uint, args map[string]interface{}) (model.Obat, error) {
	var obat model.Obat
	if err := db.DB.First(&obat, id).Error; err != nil {
		return obat, err
	}

	if v, ok := args["nama_obat"].(string); ok {
		obat.Nama_Obat = v
	}
	if v, ok := args["dosis"].(string); ok {
		obat.Dosis = v
	}
	if v, ok := args["frekuensi"].(string); ok {
		obat.Frekuensi = v
	}
	if v, ok := args["catatan"].(string); ok {
		obat.Catatan = v
	}
	if v, ok := args["tanggal"].(string); ok {
		obat.Tanggal = v
	}
	if list, ok := args["waktu"].([]interface{}); ok && len(list) > 0 {
		var waktuList []string
		for _, val := range list {
			if s, ok := val.(string); ok {
				waktuList = append(waktuList, s)
			}
		}
		obat.Waktu = strings.Join(waktuList, ",")
	}

	db.DB.Save(&obat)
	return obat, nil
}

func DeleteObat(id uint) (bool, error) {
	var obat model.Obat
	if err := db.DB.Delete(&obat, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func LoginUser(email, password string) (*model.User, error) {
	var user model.User

	// Cari user berdasarkan email
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	// Verifikasi password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("password salah")
	}

	return &user, nil
}

func ResetPassword(email, password string) (*model.User, error) {
	var user model.User

	// Cek user berdasarkan email
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	// Hash password baru
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("gagal meng-hash password")
	}

	// Simpan password baru
	user.Password = string(hashedPassword)

	if err := db.DB.Save(&user).Error; err != nil {
		return nil, errors.New("gagal menyimpan password baru")
	}

	return &user, nil
}
