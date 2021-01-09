package models

func GetStudentColumn() string {
	return `student_info.student_id",
"student_info.student_name",
"student_info.student_password",
"student_info.student_type",
"student_info.student_sex", 
"student_info.student_college",
"student_info.student_birth",
"student_info.student_time",
"student_info.student_created_time",
"student_info.student_updated_time`
}

func GetClassColumn() string {
	return `class_info.class_id",
"class_info.class_name",
"class_info.class_created_time",
"class_info.class_updated_time`
}

func GetTeacherColumn() string {
	return `teacher_info.teacher_id",
"teacher_info.teacher_name",
"teacher_info.teacher_password",
"teacher_info.teacher_type",
"teacher_info.teacher_sex", 
"teacher_info.teacher_college",
"teacher_info.teacher_birth",
"teacher_info.teacher_time",
"teacher_info.teacher_created_time",
"teacher_info.teacher_updated_time`
}

func GetCourseColumn() string {
	return `course_info.course_id",
"course_info.course_name",
"course_info.course_properties",
"course_info.course_scores",
"course_info.course_way",
"course_info.course_count", 
"course_info.course_created_time",
"course_info.course_updated_time`
}

func GetCourseBaseColumn() string {
	return `course_base_info.course_base_id",
"course_base_info.course_start_year",
"course_base_info.course_end_year",
"course_base_info.course_year",
"course_base_info.course_start_week",
"course_base_info.course_end_week",
"course_base_info.course_week",
"course_base_info.course_start_count",
"course_base_info.course_end_count",
"course_base_info.course_school",
"course_base_info.course_address",
"course_base_info.course_address_floor",
"course_base_info.course_address_number",
"course_base_info.course_base_created_time",
"course_base_info.course_base_updated_time`
}

func GetClassGroupColumn() string {
	return `class_group_info.class_group_id",
"class_group_info.class_group_name",
"class_group_info.class_group_created_time",
"class_group_info.class_group_updated_time`
}

func GetCourseStudentRelColumn() string {
	return `course_student_rel.course_student_rel_id",
"course_student_rel.student_id",
"course_student_rel.student_results",
"course_student_rel.student_point`
}

func GetClassGroupTeacherRelColumn() string {
	return `class_group_teacher_rel.is_charge`
}

func GetCourseClassRelColumn() string {
	return ""
}

func GetCourseTeacherRelColumn() string {
	return ""
}

func GetCourseGroupRelColumn() string {
	return ""
}
