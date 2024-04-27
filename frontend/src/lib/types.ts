export interface User {
    id: number;
    username: string;
    avatar: string;
    activated: boolean;
    created_at: Date;
}

export interface Course {
    id: number;
    number: string;
    name: string;
    department: string;
    units: number;
    description: string | null;
    // numRatings: number;
    // ratingRubric: RatingRubric;
    // ratingStats: RatingStats;
}

export interface Instructor {
    id: number;
    name: string;
    // teaches: CourseOffering[];
    // ratingRubric: RatingRubric;
}


export interface Offering {
    id: number;
    // course_id: number;
    course: Course;
    semester: "Fall" | "Summer 1" | "Summer 2" | "Spring" | "Winter";
    year: number;
    location: string;
    instructors: Instructor[];
}

export interface CourseOffering extends Offering {
    instructors: Instructor[];
}

export interface InstructorOffering {
    offering: Offering;
    course: Course;
}

// export interface RatingRubric {
//     overall: number;
//     teaching: number;
//     materials: number;
//     value: number;
//     difficulty: number;
//     workload: number;
//     grading: number;
// }

// Rating category => number of instances
export interface RatingStats {
    rating_count: number;
    overall: { [key: number]: number };
    teaching: { [key: number]: number };
    materials: { [key: number]: number };
    value: { [key: number]: number };
    difficulty: { [key: number]: number };
    workload: { [key: number]: number };
    grading: { [key: number]: number };
    avg_overall: number;
    avg_teaching: number;
    avg_materials: number;
    avg_value: number;
    avg_difficulty: number;
    avg_workload: number;
    avg_grading: number;
}

export interface RatingIn {
    rating_id?: number;
    offering_id: number;
    overall: number;
    teaching: number;
    materials: number;
    value: number;
    difficulty: number;
    workload: number;
    grading: number;
    comment: string;
}

export interface Rating {
    id: number;
    user: User;
    // user_id: number;
    // author: User;
    // courseOffering: CourseOffering;
    offering: Offering;
    // offering_id: number;
    created_at: Date;
    updated_at: Date;
    // rubric: RatingRubric;
    overall: number;
    teaching: number;
    materials: number;
    value: number;
    difficulty: number;
    workload: number;
    grading: number;
    comment: string;
    net_likes: number;
    liked_by_viewer: boolean;
    disliked_by_viewer: boolean;
}

// export interface RatingOutput extends Rating {
//     username: string;
//     avatar: string;
// }

// export interface Teaches {
//     id: number;
//     instructorID: number;
//     offeringID: number;
// }

export interface Metadata {
    current_page: number;
    first_page: number;
    last_page: number;
    page_size: number;
    total: number;
    totalRecords: number;
}

export type Fetch = (input: RequestInfo | URL, init?: RequestInit | undefined) => Promise<Response>

export interface UserStats {
    likes_received: number;
    courses_rated: number;
}
