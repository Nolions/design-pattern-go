package combination

import "testing"

func TestCombination(t *testing.T)  {
	f1 := NewFile("file1.txt")
	f2 := NewFile("file2.txt")
	f3 := NewFile("file3.txt")

	folder1 := NewFolder("folder1")
	folder2 := NewFolder("folder2")

	folder1.Add(&f1)
	folder1.Add(&f2)

	folder2.Add(&f3)
	folder2.Add(&folder1)
	folder2.Display()
}
