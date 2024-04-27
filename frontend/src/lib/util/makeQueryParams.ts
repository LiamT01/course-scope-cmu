interface RatingParams {
    number?: string | null;
    name?: string | null;
    courseID?: number | null;
    semester?: string | null;
    year?: string | null;
    instructorIDs?: string[] | null;
    overall?: string | null;
    sort?: string | null;
    page?: string | number | null;
    pageSize?: string | number | null;
}

export default function makeQueryParams(params: RatingParams) {
    const numberQuery = params.number ? `number=${params.number}` : null;
    const nameQuery = params.name ? `name=${params.name}` : null;
    const courseIDQuery = params.courseID ? `course_id=${params.courseID}` : null;
    const semesterQuery = params.semester && params.semester !== "all" ? `semester=${params.semester}` : null;
    const yearQuery = params.year && params.year !== "all" ? `year=${params.year}` : null;
    const instructorIDsQuery = params.instructorIDs && params.instructorIDs.length > 0 ?
        params.instructorIDs.map(id => `instructor_ids=${id}`).join("&") : null
    const overallQuery = params.overall && params.overall !== "all" ? `overall=${params.overall}` : null;

    const sortQuery = params.sort ? `sort=${params.sort}` : null;
    const pageQuery = params.page ? `page=${params.page}` : null;
    const pageSizeQuery = params.pageSize ? `page_size=${params.pageSize}` : null;

    return [numberQuery, nameQuery, courseIDQuery, semesterQuery, yearQuery, instructorIDsQuery, overallQuery, sortQuery, pageQuery, pageSizeQuery]
        .filter(Boolean)
        .join("&");
}
