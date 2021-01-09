package models

import (
"github.com/beego/beego/v2/client/orm"
	"strconv"
	"time"
)

func ParseStudentInfo(param orm.Params) *StudentInfo {
	studentId := param["student_id"]
	if studentId == nil {
		return nil
	}
	student := &StudentInfo{StudentId: studentId.(string)}
	studentName := param["student_name"]
	if studentName != nil {
		student.StudentName = studentName.(string)
	}
	studentPassword := param["student_password"]
	if studentPassword != nil {
		student.StudentPassword = studentPassword.(string)
	}
	studentType := param["student_type"]
	if studentType != nil {
		student.StudentType = studentType.(string)
	}
	studentSex := param["student_sex"]
	if studentSex != nil {
		student.StudentSex = studentSex.(string)
	}
	studentCollege := param["student_college"]
	if studentCollege != nil {
		student.StudentCollege = studentCollege.(string)
	}
	studentBirth := param["student_birth"]
	if studentBirth != nil {
		student.StudentBirth, _ = time.Parse(time.RFC3339Nano, studentBirth.(string))
	}
	studentTime := param["student_time"]
	if studentTime != nil {
		student.StudentTime, _ = time.Parse(time.RFC3339Nano, studentTime.(string))
	}
	studentResults := param["student_results"]
	if studentResults != nil {
		student.StudentResults, _ = strconv.ParseFloat(studentResults.(string), 64)
	}
	studentPoint := param["student_point"]
	if studentPoint != nil {
		student.StudentPoint, _ = strconv.ParseFloat(studentPoint.(string), 64)
	}
	studentCreatedTime := param["student_created_time"]
	if studentCreatedTime != nil {
		student.StudentCreatedTime, _ = time.Parse(time.RFC3339Nano, studentCreatedTime.(string))
	}
	studentUpdatedTime := param["student_updated_time"]
	if studentUpdatedTime != nil {
		student.StudentUpdatedTime, _ = time.Parse(time.RFC3339Nano, studentUpdatedTime.(string))
	}
	return student
}

func ParseClassInfo(param orm.Params) *ClassInfo {
	classId := param["class_id"]
	if classId == nil {
		return nil
	}
	class := &ClassInfo{ClassId: classId.(string)}
	className := param["class_name"]
	if className != nil {
		class.ClassName = className.(string)
	}
	classCreatedTime := param["class_created_time"]
	if classCreatedTime != nil {
		class.ClassCreatedTime, _ = time.Parse(time.RFC3339Nano, classCreatedTime.(string))
	}
	classUpdatedTime := param["class_updated_time"]
	if classUpdatedTime != nil {
		class.ClassUpdatedTime, _ = time.Parse(time.RFC3339Nano, classUpdatedTime.(string))
	}
	return class
}

func ParseTeacherInfo(param orm.Params) *TeacherInfo {
	teacherId := param["teacher_id"]
	if teacherId == nil {
		return nil
	}
	teacher := &TeacherInfo{TeacherId: teacherId.(string)}
	teacherName := param["teacher_name"]
	if teacherName != nil {
		teacher.TeacherName = teacherName.(string)
	}
	teacherPassword := param["teacher_password"]
	if teacherPassword != nil {
		teacher.TeacherPassword = teacherPassword.(string)
	}
	teacherType := param["teacher_type"]
	if teacherType != nil {
		teacher.TeacherType = teacherType.(string)
	}
	teacherSex := param["teacher_sex"]
	if teacherSex != nil {
		teacher.TeacherSex = teacherSex.(string)
	}
	teacherCollege := param["teacher_college"]
	if teacherCollege != nil {
		teacher.TeacherCollege = teacherCollege.(string)
	}
	teacherBirth := param["teacher_birth"]
	if teacherBirth != nil {
		teacher.TeacherBirth, _ = time.Parse(time.RFC3339Nano, teacherBirth.(string))
	}
	teacherTime := param["teacher_time"]
	if teacherTime != nil {
		teacher.TeacherTime, _ = time.Parse(time.RFC3339Nano, teacherTime.(string))
	}
	isCharge := param["is_charge"]
	if isCharge != nil {
		teacher.IsCharge = isCharge.(string) == "true"
	}
	teacherCreatedTime := param["teacher_created_time"]
	if teacherCreatedTime != nil {
		teacher.TeacherCreatedTime, _ = time.Parse(time.RFC3339Nano, teacherCreatedTime.(string))
	}
	teacherUpdatedTime := param["teacher_updated_time"]
	if teacherUpdatedTime != nil {
		teacher.TeacherUpdatedTime, _ = time.Parse(time.RFC3339Nano, teacherUpdatedTime.(string))
	}
	return teacher
}

func ParseCourseInfo(param orm.Params) *CourseInfo {
	courseId := param["course_id"]
	if courseId == nil {
		return nil
	}
	course := &CourseInfo{CourseId: courseId.(string)}
	courseName := param["course_name"]
	if courseName != nil {
		course.CourseName = courseName.(string)
	}
	courseProperties := param["course_properties"]
	if courseProperties != nil {
		course.CourseProperties = courseProperties.(string)
	}
	courseScores := param["course_scores"]
	if courseScores != nil {
		course.CourseScores, _ = strconv.ParseFloat(courseScores.(string), 64)
	}
	courseWay := param["course_way"]
	if courseWay != nil {
		course.CourseWay = courseWay.(string)
	}
	courseCount := param["course_count"]
	if courseCount != nil {
		course.CourseCount, _ = strconv.ParseFloat(courseCount.(string), 64)
	}
	studentResults := param["student_results"]
	if studentResults != nil {
		course.StudentResults, _ = strconv.ParseFloat(studentResults.(string), 64)
	}
	studentPoint := param["student_point"]
	if studentPoint != nil {
		course.StudentPoint, _ = strconv.ParseFloat(studentPoint.(string), 64)
	}
	courseCreatedTime := param["course_created_time"]
	if courseCreatedTime != nil {
		course.CourseCreatedTime, _ = time.Parse(time.RFC3339Nano, courseCreatedTime.(string))
	}
	courseUpdatedTime := param["course_updated_time"]
	if courseUpdatedTime != nil {
		course.CourseUpdatedTime, _ = time.Parse(time.RFC3339Nano, courseUpdatedTime.(string))
	}
	return course
}

func ParseCourseBaseInfo(param orm.Params) *CourseBaseInfo {
	courseBaseIdStr := param["course_base_id"]
	if courseBaseIdStr == nil {
		return nil
	}
	courseBaseId, _ := strconv.Atoi(courseBaseIdStr.(string))
	courseBase := &CourseBaseInfo{CourseBaseId: courseBaseId}
	courseStartYear := param["course_start_year"]
	if courseStartYear != nil {
		courseBase.CourseStartYear, _ = strconv.Atoi(courseStartYear.(string))
	}
	courseEndYear := param["course_end_year"]
	if courseEndYear != nil {
		courseBase.CourseEndYear, _ = strconv.Atoi(courseEndYear.(string))
	}
	courseYear := param["course_year"]
	if courseYear != nil {
		courseBase.CourseYear, _ = strconv.Atoi(courseYear.(string))
	}
	courseStartWeek := param["course_start_week"]
	if courseStartWeek != nil {
		courseBase.CourseStartWeek, _ = strconv.Atoi(courseStartWeek.(string))
	}
	courseEndWeek := param["course_end_week"]
	if courseEndWeek != nil {
		courseBase.CourseEndWeek, _ = strconv.Atoi(courseEndWeek.(string))
	}
	courseWeek := param["course_week"]
	if courseWeek != nil {
		courseBase.CourseWeek, _ = strconv.Atoi(courseWeek.(string))
	}
	courseStartCount := param["course_start_count"]
	if courseStartCount != nil {
		courseBase.CourseStartCount, _ = strconv.Atoi(courseStartCount.(string))
	}
	courseEndCount := param["course_end_count"]
	if courseEndCount != nil {
		courseBase.CourseEndCount, _ = strconv.Atoi(courseEndCount.(string))
	}
	courseSchool := param["course_school"]
	if courseSchool != nil {
		courseBase.CourseSchool = courseSchool.(string)
	}
	courseAddress := param["course_address"]
	if courseAddress != nil {
		courseBase.CourseAddress = courseAddress.(string)
	}
	courseAddressFloor := param["course_address_floor"]
	if courseAddressFloor != nil {
		courseBase.CourseAddressFloor, _ = strconv.Atoi(courseAddressFloor.(string))
	}
	courseAddressNumber := param["course_address_number"]
	if courseAddressNumber != nil {
		courseBase.CourseAddressNumber, _ = strconv.Atoi(courseAddressNumber.(string))
	}
	courseBaseCreatedTime := param["course_base_created_time"]
	if courseBaseCreatedTime != nil {
		courseBase.CourseBaseCreatedTime, _ = time.Parse(time.RFC3339Nano, courseBaseCreatedTime.(string))
	}
	courseBaseUpdatedTime := param["course_base_updated_time"]
	if courseBaseUpdatedTime != nil {
		courseBase.CourseBaseUpdatedTime, _ = time.Parse(time.RFC3339Nano, courseBaseUpdatedTime.(string))
	}
	return courseBase
}

func ParseClassGroupInfo(param orm.Params) *ClassGroupInfo {
	classGroupIdStr := param["class_group_id"]
	if classGroupIdStr == nil {
		return nil
	}
	classGroupId, _ := strconv.Atoi(classGroupIdStr.(string))
	classGroup := &ClassGroupInfo{ClassGroupId: classGroupId}
	classGroupName := param["class_group_name"]
	if classGroupName != nil {
		classGroup.ClassGroupName = classGroupName.(string)
	}
	isCharge := param["is_charge"]
	if isCharge != nil {
		classGroup.IsCharge = isCharge.(string) == "true"
	}
	classGroupCreatedTime := param["class_group_created_time"]
	if classGroupCreatedTime != nil {
		classGroup.ClassGroupCreatedTime, _ = time.Parse(time.RFC3339Nano, classGroupCreatedTime.(string))
	}
	classGroupUpdatedTime := param["class_group_updated_time"]
	if classGroupUpdatedTime != nil {
		classGroup.ClassGroupUpdatedTime, _ = time.Parse(time.RFC3339Nano, classGroupUpdatedTime.(string))
	}
	return classGroup
}

func ParseCourseStudentRel(param orm.Params) *CourseStudentRel {
	return nil
}

func ParseClassGroupTeacherRel(param orm.Params) *ClassGroupTeacherRel {
	return nil
}

func ParseCourseClassRel(param orm.Params) *CourseClassRel {
	return nil
}

func ParseCourseTeacherRel(param orm.Params) *CourseTeacherRel {
	return nil
}

func ParseCourseGroupRel(param orm.Params) *CourseGroupRel {
	return nil
}

func ParseCourses(params []orm.Params) []*CourseInfo {
	var courses []*CourseInfo
	courseMap := make(map[string]*CourseInfo)
	studentMap := make(map[string]map[string]*StudentInfo)
	classMap := make(map[string]map[string]*ClassInfo)
	teacherMap := make(map[string]map[string]*TeacherInfo)
	classGroupMap := make(map[string]map[int]*ClassGroupInfo)
	for _, param := range params {
		course := ParseCourseInfo(param)
		if course == nil {
			continue
		}
		student := ParseStudentInfo(param)
		class := ParseClassInfo(param)
		teacher := ParseTeacherInfo(param)
		classGroup := ParseClassGroupInfo(param)
		courseBase := ParseCourseBaseInfo(param)
		if oldCourse, ok := courseMap[course.CourseId]; ok {
			if student != nil {
				if _, ok := studentMap[course.CourseId][student.StudentId]; !ok {
					oldCourse.Students = append(oldCourse.Students, student)
					continue
				}
			}
			if class != nil {
				if _, ok := classMap[course.CourseId][class.ClassId]; !ok {
					oldCourse.Classes = append(oldCourse.Classes, class)
					continue
				}
			}
			if teacher != nil {
				if _, ok := teacherMap[course.CourseId][teacher.TeacherId]; !ok {
					oldCourse.Teachers = append(oldCourse.Teachers, teacher)
					continue
				}
			}
			if classGroup != nil {
				if _, ok := classGroupMap[course.CourseId][classGroup.ClassGroupId]; !ok {
					oldCourse.ClassGroups = append(oldCourse.ClassGroups, classGroup)
					continue
				}
			}
			oldCourse.CourseBases = append(oldCourse.CourseBases, courseBase)
			continue
		}
		if _, ok := studentMap[course.CourseId]; !ok {
			studentMap[course.CourseId] = make(map[string]*StudentInfo)
		}
		if _, ok := classMap[course.CourseId]; !ok {
			classMap[course.CourseId] = make(map[string]*ClassInfo)
		}
		if _, ok := teacherMap[course.CourseId]; !ok {
			teacherMap[course.CourseId] = make(map[string]*TeacherInfo)
		}
		if _, ok := classGroupMap[course.CourseId]; !ok {
			classGroupMap[course.CourseId] = make(map[int]*ClassGroupInfo)
		}
		courseMap[course.CourseId] = course

		if student != nil {
			studentMap[course.CourseId][student.StudentId] = student
			course.Students = append(course.Students, student)
		}
		if class != nil {
			classMap[course.CourseId][class.ClassId] = class
			course.Classes = append(course.Classes, class)
		}
		if teacher != nil {
			teacherMap[course.CourseId][teacher.TeacherId] = teacher
			course.Teachers = append(course.Teachers, teacher)
		}
		if classGroup != nil {
			classGroupMap[course.CourseId][classGroup.ClassGroupId] = classGroup
			course.ClassGroups = append(course.ClassGroups, classGroup)
		}
		if courseBase != nil {
			course.CourseBases = append(course.CourseBases, courseBase)
		}

		courses = append(courses, course)
	}
	return courses
}

func ParseCourseStudent(params []orm.Params) []*StudentInfo {
	var students []*StudentInfo
	for _, param := range params {
		student := ParseStudentInfo(param)
		if student == nil {
			continue
		}
		class := ParseClassInfo(param)
		student.Class = class
		students = append(students, student)
	}
	return students
}
