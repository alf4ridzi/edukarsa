package domain

import "errors"

var (
	// error auth
	ErrUsernameExist = errors.New("username already exist")
	ErrEmailExist    = errors.New("email already exist")
	ErrWrongPassword = errors.New("password is wrong")
	// error common
	ErrForbidden = errors.New("forbidden")
	// error class
	ErrAlreadyJoinedClass  = errors.New("sudah bergabung")
	ErrNotJoinedClass      = errors.New("belum bergabung")
	ErrCreatorCantLeave    = errors.New("kreator tidak boleh keluar")
	ErrInvalidExtension    = errors.New("extensi file tidak disupport")
	ErrFileSizeTooBig      = errors.New("ukuran file terlalu besar")
	ErrMinimumOption       = errors.New("minimal pilihan adalah 2")
	ErrInvalidCorrectIndex = errors.New("jawaban opsi tidak benar")
	// error exam
	ErrExamNotStarted       = errors.New("ujian belum dimulai")
	ErrExamAlreadyFinished  = errors.New("ujian sudah selesai")
	ErrExamNotAccessible    = errors.New("ujian belum dapat diakses")
	ErrAlreadyStartExam     = errors.New("sudah memulai ujian")
	ErrUserExamNotStarted   = errors.New("ujian belum dimulai oleh user")
	ErrExamAlreadySubmitted = errors.New("ujian sudah disubmit")
	ErrExamExpired          = errors.New("ujian expired")
	// error submission
	ErrInvalidSubmissionStatus = errors.New("invalid status submit")
	ErrSubmissionNotStarted    = errors.New("tidak dapat submit karena belum dimulai")
	// error question
	ErrQuestionNotBelongToExam = errors.New("pertanyaan tidak valid dengan ujian")
	// error option
	ErrOptionNotBelongToQuestion = errors.New("opsi tidak valid dengan pertanyaan")
	// error answer
	ErrSameAnswerSubmitted = errors.New("opsi tidak boleh sama")
)
