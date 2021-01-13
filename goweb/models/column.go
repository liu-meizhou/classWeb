package models

func GetStudentColumn() string {
	return `student_info.student_id",
"student_info.student_name",
"student_info.student_password",
"student_info.student_type",
"student_info.student_sex", 
"student_info.student_college",
"student_info.student_all_point",
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
	// "course_info.course_max_number",
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

func GetCourseGroupColumn() string {
	return `course_group_info.course_group_id",
"course_group_info.course_group_name",
"course_group_info.course_group_created_time",
"course_group_info.course_group_updated_time`
}

func GetCourseStudentRelColumn() string {
	return `course_student_rel.course_student_rel_id",
"course_student_rel.student_id",
"course_student_rel.course_id",
"course_student_rel.student_results",
"course_student_rel.student_point`
}

func GetCourseGroupTeacherRelColumn() string {
	return `course_group_teacher_rel.course_group_teacher_rel_id",
"course_group_teacher_rel.course_group_id",
"course_group_teacher_rel.teacher_id",
"course_group_teacher_rel.is_charge`
}

func GetCourseClassRelColumn() string {
	return `course_class_rel.course_class_rel_id",
"course_class_rel.course_id",
"course_class_rel.class_id`
}

func GetCourseTeacherRelColumn() string {
	return `course_teacher_rel.course_teacher_rel_id",
"course_teacher_rel.course_id",
"course_teacher_rel.teacher_id`
}

func GetCourseGroupRelColumn() string {
	return `course_group_rel.course_group_rel_id",
"course_group_rel.course_id",
"course_group_rel.course_group_id`
}
