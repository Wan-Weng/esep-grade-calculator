package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "A" //change to A

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 95, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 91, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestEmptyCategoriesAreZero(t *testing.T) {
	gc := NewGradeCalculator()
	if got := gc.calculateNumericalGrade(); got != 0 {
		t.Fatalf("empty categories: numeric=%v, want 0", got)
	}
}

func TestGetGradeC(t *testing.T) {
	gc := NewGradeCalculator()

	gc.AddGrade("a", 70, Assignment)
	gc.AddGrade("e", 70, Exam)
	gc.AddGrade("s", 70, Essay)

	if got := gc.GetFinalGrade(); got != "C" {
		t.Fatalf("want C, got %s", got)
	}
}

func TestGetGradeD(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("a", 60, Assignment)
	gc.AddGrade("e", 60, Exam)
	gc.AddGrade("s", 60, Essay)

	if got := gc.GetFinalGrade(); got != "D" {
		t.Fatalf("want D, got %s", got)
	}
}

func TestGetGradeF_Empty(t *testing.T) {
	gc := NewGradeCalculator() // all categories empty → numeric = 0
	if got := gc.GetFinalGrade(); got != "F" {
		t.Fatalf("want F, got %s", got)
	}
}

func TestGradeTypeString(t *testing.T) {
	cases := []struct {
		in   GradeType
		want string
	}{
		{Assignment, "assignment"},
		{Exam, "exam"},
		{Essay, "essay"},
	}
	for i := range cases {
		got := cases[i].in.String()
		if got != cases[i].want {
			t.Fatalf("case %d (%v): got %q, want %q", i, cases[i].in, got, cases[i].want)
		}
	}
}
