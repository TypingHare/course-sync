package service_test

import (
	"errors"
	"testing"
	"time"

	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/domain/service"
)

func TestStudentServiceAddStudent(t *testing.T) {
	t.Parallel()

	repo := &studentRepoStub{
		getAllResult: []model.Student{{ID: 1, Name: "alice"}},
	}
	svc := service.NewStudentService(repo)

	if err := svc.AddStudent(&model.Student{ID: 2, Name: "bob"}); err != nil {
		t.Fatalf("AddStudent returned error: %v", err)
	}

	if repo.saveAllCalls != 1 {
		t.Fatalf("SaveAll call count = %d, want 1", repo.saveAllCalls)
	}
	if len(repo.saved) != 2 {
		t.Fatalf("saved length = %d, want 2", len(repo.saved))
	}
	if repo.saved[1].Name != "bob" {
		t.Fatalf("saved second student = %q, want %q", repo.saved[1].Name, "bob")
	}
}

func TestStudentServiceGetNextStudentID(t *testing.T) {
	t.Parallel()

	repo := &studentRepoStub{
		getAllResult: []model.Student{{ID: 3}, {ID: 9}, {ID: 2}},
	}
	svc := service.NewStudentService(repo)

	got, err := svc.GetNextStudentID()
	if err != nil {
		t.Fatalf("GetNextStudentID returned error: %v", err)
	}
	if got != 10 {
		t.Fatalf("GetNextStudentID = %d, want 10", got)
	}
}

func TestAssignmentServiceAddAssignmentRejectsDuplicateName(t *testing.T) {
	t.Parallel()

	repo := &assignmentRepoStub{
		getAllResult: []model.Assignment{{Name: "hw1"}},
	}
	svc := service.NewAssignmentService(repo)

	err := svc.AddAssignment(&model.Assignment{Name: "hw1"})
	if err == nil {
		t.Fatalf("AddAssignment succeeded for duplicate name; want error")
	}
	if repo.saveAllCalls != 0 {
		t.Fatalf("SaveAll call count = %d, want 0", repo.saveAllCalls)
	}
}

func TestAssignmentServiceGetAssignmentByName(t *testing.T) {
	t.Parallel()

	repo := &assignmentRepoStub{
		getAllResult: []model.Assignment{
			{Name: "hw1"},
			{Name: "hw2", Title: "Homework 2"},
		},
	}
	svc := service.NewAssignmentService(repo)

	got, err := svc.GetAssignmentByName("hw2")
	if err != nil {
		t.Fatalf("GetAssignmentByName returned error: %v", err)
	}
	if got == nil {
		t.Fatalf("GetAssignmentByName returned nil, want assignment")
	}
	if got.Title != "Homework 2" {
		t.Fatalf("assignment title = %q, want %q", got.Title, "Homework 2")
	}
}

func TestSubmissionServiceAddSubmission(t *testing.T) {
	t.Parallel()

	repo := &submissionRepoStub{
		getAllResult: []model.Submission{{Hash: "a"}},
	}
	svc := service.NewSubmissionService(repo)

	if err := svc.AddSubmission(&model.Submission{Hash: "b"}); err != nil {
		t.Fatalf("AddSubmission returned error: %v", err)
	}
	if repo.saveAllCalls != 1 {
		t.Fatalf("SaveAll call count = %d, want 1", repo.saveAllCalls)
	}
	if len(repo.saved) != 2 {
		t.Fatalf("saved length = %d, want 2", len(repo.saved))
	}
}

func TestGradeServiceGetLastGradeByAssignmentName(t *testing.T) {
	t.Parallel()

	oldTime := time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC)
	newTime := time.Date(2025, 1, 2, 10, 0, 0, 0, time.UTC)

	repo := &gradeRepoStub{
		getAllResult: []model.Grade{
			{AssignmentName: "hw1", Score: 80, GradedAt: oldTime},
			{AssignmentName: "hw2", Score: 90, GradedAt: newTime},
			{AssignmentName: "hw1", Score: 95, GradedAt: newTime},
		},
	}
	svc := service.NewGradeService(repo)

	got, err := svc.GetLastGradeByAssignmentName("hw1")
	if err != nil {
		t.Fatalf("GetLastGradeByAssignmentName returned error: %v", err)
	}
	if got == nil {
		t.Fatalf("GetLastGradeByAssignmentName returned nil, want grade")
	}
	if got.Score != 95 {
		t.Fatalf("latest score = %.2f, want 95", got.Score)
	}
}

func TestDocServiceGetDefaultDoc(t *testing.T) {
	t.Parallel()

	repo := &docRepoStub{
		getAllResult: []model.Doc{
			{Name: "guide-v1", IsDefault: false},
			{Name: "guide-v2", IsDefault: true},
		},
	}
	svc := service.NewDocService(repo)

	got, err := svc.GetDefaultDoc()
	if err != nil {
		t.Fatalf("GetDefaultDoc returned error: %v", err)
	}
	if got == nil {
		t.Fatalf("GetDefaultDoc returned nil, want doc")
	}
	if got.Name != "guide-v2" {
		t.Fatalf("default doc = %q, want %q", got.Name, "guide-v2")
	}
}

func TestServicesPropagateRepoGetAllErrors(t *testing.T) {
	t.Parallel()

	getAllErr := errors.New("repo get all failure")

	if err := service.NewStudentService(&studentRepoStub{getAllErr: getAllErr}).AddStudent(&model.Student{}); err == nil {
		t.Fatalf("StudentService.AddStudent should propagate GetAll error")
	}
	if _, err := service.NewAssignmentService(&assignmentRepoStub{getAllErr: getAllErr}).GetAssignmentByName("x"); err == nil {
		t.Fatalf("AssignmentService.GetAssignmentByName should propagate GetAll error")
	}
	if err := service.NewSubmissionService(&submissionRepoStub{getAllErr: getAllErr}).AddSubmission(&model.Submission{}); err == nil {
		t.Fatalf("SubmissionService.AddSubmission should propagate GetAll error")
	}
	if _, err := service.NewGradeService(&gradeRepoStub{getAllErr: getAllErr}).GetGradeBySubmissionHash("x"); err == nil {
		t.Fatalf("GradeService.GetGradeBySubmissionHash should propagate GetAll error")
	}
	if _, err := service.NewDocService(&docRepoStub{getAllErr: getAllErr}).GetDocByName("x"); err == nil {
		t.Fatalf("DocService.GetDocByName should propagate GetAll error")
	}
}

type studentRepoStub struct {
	getAllResult []model.Student
	getAllErr    error
	saved        []model.Student
	saveAllCalls int
}

func (s *studentRepoStub) GetAll() ([]model.Student, error) {
	return s.getAllResult, s.getAllErr
}

func (s *studentRepoStub) SaveAll(entities []model.Student) error {
	s.saved = append([]model.Student(nil), entities...)
	s.saveAllCalls++
	return nil
}

type assignmentRepoStub struct {
	getAllResult []model.Assignment
	getAllErr    error
	saved        []model.Assignment
	saveAllCalls int
}

func (s *assignmentRepoStub) GetAll() ([]model.Assignment, error) {
	return s.getAllResult, s.getAllErr
}

func (s *assignmentRepoStub) SaveAll(entities []model.Assignment) error {
	s.saved = append([]model.Assignment(nil), entities...)
	s.saveAllCalls++
	return nil
}

type submissionRepoStub struct {
	getAllResult []model.Submission
	getAllErr    error
	saved        []model.Submission
	saveAllCalls int
}

func (s *submissionRepoStub) GetAll() ([]model.Submission, error) {
	return s.getAllResult, s.getAllErr
}

func (s *submissionRepoStub) SaveAll(entities []model.Submission) error {
	s.saved = append([]model.Submission(nil), entities...)
	s.saveAllCalls++
	return nil
}

type gradeRepoStub struct {
	getAllResult []model.Grade
	getAllErr    error
}

func (s *gradeRepoStub) GetAll() ([]model.Grade, error) {
	return s.getAllResult, s.getAllErr
}

func (s *gradeRepoStub) SaveAll([]model.Grade) error {
	return nil
}

type docRepoStub struct {
	getAllResult []model.Doc
	getAllErr    error
}

func (s *docRepoStub) GetAll() ([]model.Doc, error) {
	return s.getAllResult, s.getAllErr
}

func (s *docRepoStub) SaveAll([]model.Doc) error {
	return nil
}
