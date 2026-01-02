package domain

import "errors"

var (
	ErrUsernameExist             = errors.New("username already exist")
	ErrEmailExist                = errors.New("email already exist")
	ErrWrongPassword             = errors.New("password is wrong")
	ErrForbidden                 = errors.New("forbidden")
	ErrAlreadyJoinedClass        = errors.New("sudah bergabung")
	ErrNotJoinedClass            = errors.New("belum bergabung")
	ErrCreatorCantLeave          = errors.New("kreator tidak boleh keluar")
	ErrInvalidExtension          = errors.New("extensi file tidak disupport")
	ErrFileSizeTooBig            = errors.New("ukuran file terlalu besar")
	ErrMinimumOption             = errors.New("minimal pilihan adalah 2")
	ErrInvalidCorrectIndex       = errors.New("jawaban opsi tidak benar")
	ErrExamNotStarted            = errors.New("ujian belum dimulai")
	ErrExamAlreadyFinished       = errors.New("ujian sudah selesai")
	ErrExamNotAccessible         = errors.New("ujian belum dapat diakses")
	ErrQuestionNotBelongToExam   = errors.New("pertanyaan tidak valid dengan ujian")
	ErrOptionNotBelongToQuestion = errors.New("opsi tidak valid dengan pertanyaan")
	ErrSameAnswerSubmitted       = errors.New("opsi tidak boleh sama")
)
