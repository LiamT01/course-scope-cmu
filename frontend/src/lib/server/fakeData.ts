// import {faker} from "@faker-js/faker";
// import type {
//     Course,
//     CourseOffering,
//     DifficultyLevel,
//     GradingLevel,
//     Instructor,
//     Rating,
//     RatingScale,
//     User,
//     ValidLevelType,
//     WorkloadLevel
// } from "$lib/types";
//
// const difficultyScale: Record<DifficultyLevel, number> = {
//     "effortless": 1, "manageable": 2, "standard": 3, "tough": 4, "brutal": 5
// }
//
// const workloadScale: Record<WorkloadLevel, number> = {
//     "light": 1, "moderate": 2, "balanced": 3, "heavy": 4, "overwhelming": 5
// }
//
// const gradingScale: Record<GradingLevel, number> = {
//     "lenient": 1, "flexible": 2, "fair": 3, "strict": 4, "rigorous": 5
// }
//
// const calcAvgTextScale = <LevelType extends ValidLevelType>(ratings: LevelType[], scale: Record<LevelType, number>): LevelType => {
//     const sum: number = ratings.reduce((acc, curr) => acc + scale[curr], 0);
//     const avg: number = sum / ratings.length;
//     const roundedAvg: number = Math.round(avg);
//     return Object.keys(scale).find(key => scale[key as LevelType] === roundedAvg) as LevelType;
// }
//
// export const courses: Course[] = [{
//     id: 1,
//     number: "10-701",
//     name: "Introduction to Machine Learning (PhD)",
//     department: "Machine Learning",
//     units: 12,
//     description: 'Machine learning studies the question: "how can we build adaptive algorithms that automatically improve their performance (on a given task) as they acquire more experience?" This can cover a dizzying array of technologies depending on what sort of task we have in mind, and we take to constitute experience. Through this framing, we might view classical statistics problems, like estimating the likelihood that a coin lands on heads as an ML problem: the task is to produce an estimate, and the experience would consist of observations. But ML can also include robotics challenges, where the experience is acquired dynamically as our artificial agent interacts with the real world. Other grand challenges in machine learning relate to personalized medicine, natural language processing, and most recently generating media artifacts like photographs and essays (but don\'t ask chatGPT to do your homework). This course is designed to give PhD students a solid foundation in the methods, mathematics, and algorithms of modern machine learning. Students entering the class with a pre-existing working knowledge of probability, statistics and algorithms will be at an advantage, but the class has been designed so that anyone with a strong mathematical and computer science background can catch up and fully participate. If you are interested in this topic, but are not a PhD student, or are a PhD student not specializing in machine learning, you might consider the master\'s level course on Machine Learning, 10-601. This class may be appropriate for MS and undergrad students who are interested in the theory and algorithms behind ML.',
//     numRatings: 0,
//     ratingRubric: {
//         overallRating: 0,
//         teachingQuality: 0,
//         materialsQuality: 0,
//         amountGains: 0,
//         difficulty: "standard",
//         workload: "balanced",
//         grading: "fair",
//     },
//     ratingStats: {
//         overallRating: [],
//         teachingQuality: [],
//         materialsQuality: [],
//         amountGains: [],
//         difficulty: [],
//         workload: [],
//         grading: [],
//     },
// }, {
//     id: 2,
//     number: "36-700",
//     name: "Probability and Mathematical Statistics",
//     department: "Statistics",
//     units: 12,
//     description: 'This is a one-semester course covering the basics of statistics. We will first provide a quick introduction to probability theory, and then cover fundamental topics in mathematical statistics such as point estimation, hypothesis testing, asymptotic theory, and Bayesian inference. If time permits, we will also cover more advanced and useful topics including nonparametric inference, regression and classification. Prerequisites: one- and two-variable calculus and matrix algebra.  Graduate students in degree-seeking programs are given priority.',
//     numRatings: 0,
//     ratingRubric: {
//         overallRating: 0,
//         teachingQuality: 0,
//         materialsQuality: 0,
//         amountGains: 0,
//         difficulty: "standard",
//         workload: "balanced",
//         grading: "fair",
//     },
//     ratingStats: {
//         overallRating: [],
//         teachingQuality: [],
//         materialsQuality: [],
//         amountGains: [],
//         difficulty: [],
//         workload: [],
//         grading: [],
//     },
// }, {
//     id: 3,
//     number: "15-640",
//     name: "Distributed Systems",
//     department: "Computer Science",
//     units: 12,
//     description: 'The goals of this course are twofold: First, for students to gain an understanding of the principles and techniques behind the design of distributed systems, such as locking, concurrency, scheduling, and communication across the network. Second, for students to gain practical experience designing, implementing, and debugging real distributed systems.  The major themes this course will teach include scarcity, scheduling, concurrency and concurrent programming, naming, abstraction and modularity, imperfect communication and other types of failure, protection from accidental and malicious harm, optimism, and the use of instrumentation and monitoring and debugging tools in problem solving. As the creation and management of software systems is a fundamental goal of any undergraduate systems course, students will design, implement, and debug large programming projects.   As a consequence, competency in both the C and Java programming languages is required.',
//     numRatings: 0,
//     ratingRubric: {
//         overallRating: 0,
//         teachingQuality: 0,
//         materialsQuality: 0,
//         amountGains: 0,
//         difficulty: "standard",
//         workload: "balanced",
//         grading: "fair",
//     },
//     ratingStats: {
//         overallRating: [],
//         teachingQuality: [],
//         materialsQuality: [],
//         amountGains: [],
//         difficulty: [],
//         workload: [],
//         grading: [],
//     },
// },];
//
// export const instructors: Instructor[] = [{
//     id: 1,
//     name: "Chai, Henry",
//     teaches: [],
//     ratingRubric: {
//         overallRating: 0,
//         teachingQuality: 0,
//         materialsQuality: 0,
//         amountGains: 0,
//         difficulty: "standard",
//         workload: "balanced",
//         grading: "fair"
//     }
// }, {
//     id: 2,
//     name: "Lipton, Zachary",
//     teaches: [],
//     ratingRubric: {
//         overallRating: 0,
//         teachingQuality: 0,
//         materialsQuality: 0,
//         amountGains: 0,
//         difficulty: "standard",
//         workload: "balanced",
//         grading: "fair"
//     }
// }, {
//     id: 3,
//     name: "Wasserman, Larry",
//     teaches: [],
//     ratingRubric: {
//         overallRating: 0,
//         teachingQuality: 0,
//         materialsQuality: 0,
//         amountGains: 0,
//         difficulty: "standard",
//         workload: "balanced",
//         grading: "fair"
//     }
// }, {
//     id: 4,
//     name: "Miller, Heather",
//     teaches: [],
//     ratingRubric: {
//         overallRating: 0,
//         teachingQuality: 0,
//         materialsQuality: 0,
//         amountGains: 0,
//         difficulty: "standard",
//         workload: "balanced",
//         grading: "fair"
//     }
// }, {
//     id: 5,
//     name: "Zheng, Wenting",
//     teaches: [],
//     ratingRubric: {
//         overallRating: 0,
//         teachingQuality: 0,
//         materialsQuality: 0,
//         amountGains: 0,
//         difficulty: "standard",
//         workload: "balanced",
//         grading: "fair"
//     }
// }];
//
// export const courseOfferings: CourseOffering[] = [{
//     id: 1,
//     course: courses[0],
//     semester: "Fall",
//     year: 2023,
//     location: "Pittsburgh, Pennsylvania",
//     taughtBy: [instructors[0], instructors[1]],
// }, {
//     id: 2,
//     course: courses[1],
//     semester: "Fall",
//     year: 2023,
//     location: "Pittsburgh, Pennsylvania",
//     taughtBy: [instructors[2]],
// }, {
//     id: 3,
//     course: courses[2],
//     semester: "Fall",
//     year: 2023,
//     location: "Pittsburgh, Pennsylvania",
//     taughtBy: [instructors[3], instructors[4]],
// }];
//
// instructors[0].teaches.push(courseOfferings[0]);
// instructors[1].teaches.push(courseOfferings[0]);
// instructors[2].teaches.push(courseOfferings[1]);
// instructors[3].teaches.push(courseOfferings[2]);
// instructors[4].teaches.push(courseOfferings[2]);
//
// export const users: User[] = [{
//     id: 1, nickname: "Superb Fennel", createdAt: faker.date.past(),
// }, {
//     id: 2, nickname: "Lucky Cabbage", createdAt: faker.date.past(),
// }, {
//     id: 3, nickname: "Tall Huckleberry", createdAt: faker.date.past(),
// },];
//
// // Function to generate a random element from an array
// function randomElement<T>(array: T[]): T {
//     return array[Math.floor(Math.random() * array.length)];
// }
//
// export const currentUser: User = users[0];
//
// // Function to generate random ratings
// function generateRatings(users: User[], courseOfferings: CourseOffering[]): Rating[] {
//     const ratings: Rating[] = [];
//     let idCounter: number = 1;
//
//     users.forEach(user => {
//         courseOfferings.forEach(offering => {
//             const rating: Rating = {
//                 id: idCounter++,
//                 author: user,
//                 courseOffering: offering,
//                 createdAt: faker.date.past(),
//                 lastModified: faker.date.recent(),
//                 ratingRubric: {
//                     overallRating: randomElement<RatingScale>([1, 2, 3, 4, 5]),
//                     teachingQuality: randomElement<RatingScale>([1, 2, 3, 4, 5]),
//                     materialsQuality: randomElement<RatingScale>([1, 2, 3, 4, 5]),
//                     amountGains: randomElement<RatingScale>([1, 2, 3, 4, 5]),
//                     difficulty: randomElement(["effortless", "manageable", "standard", "tough", "brutal"]),
//                     workload: randomElement(["light", "moderate", "balanced", "heavy", "overwhelming"]),
//                     grading: randomElement(["lenient", "flexible", "fair", "strict", "rigorous"]),
//                 },
//                 comment: faker.lorem.paragraphs({min: 5, max: 8}),
//                 likes: faker.number.int({min: 0, max: 100}),
//                 dislikes: faker.number.int({min: 0, max: 100}),
//                 ownedByUser: user.id === currentUser.id,
//                 likedByUser: faker.datatype.boolean(),
//                 dislikedByUser: faker.datatype.boolean()
//             };
//             ratings.push(rating);
//         });
//     });
//
//     return ratings;
// }
//
// export const ratings = generateRatings(users, courseOfferings);
//
// // Calculate average rating for each course
// courses.forEach(course => {
//     const ratingsForCourse: Rating[] = ratings.filter(rating => rating.courseOffering.course.id === course.id);
//
//     course.numRatings = ratingsForCourse.length;
//
//     course.ratingRubric.overallRating = ratingsForCourse.reduce((acc, curr) => acc + curr.ratingRubric.overallRating, 0) / ratingsForCourse.length;
//     course.ratingRubric.overallRating = Math.round(course.ratingRubric.overallRating * 10) / 10;
//     let overallRatingStats: { [key: number]: number } = {1: 0, 2: 0, 3: 0, 4: 0, 5: 0};
//     ratingsForCourse.forEach(rating => {
//         overallRatingStats[rating.ratingRubric.overallRating]++;
//     });
//     for (const rating in overallRatingStats) {
//         overallRatingStats[rating] = Math.round(overallRatingStats[rating] / ratingsForCourse.length * 100);
//     }
//     course.ratingStats.overallRating = Object.keys(overallRatingStats).map(key => ({
//         label: `${key}-star`, percentage: overallRatingStats[Number(key)]
//     }));
//
//     course.ratingRubric.teachingQuality = ratingsForCourse.reduce((acc, curr) => acc + curr.ratingRubric.teachingQuality, 0) / ratingsForCourse.length;
//     course.ratingRubric.teachingQuality = Math.round(course.ratingRubric.teachingQuality * 10) / 10;
//     let teachingQualityStats: { [key: number]: number } = {1: 0, 2: 0, 3: 0, 4: 0, 5: 0};
//     ratingsForCourse.forEach(rating => {
//         teachingQualityStats[rating.ratingRubric.teachingQuality]++;
//     });
//     for (const rating in teachingQualityStats) {
//         teachingQualityStats[rating] = Math.round(teachingQualityStats[rating] / ratingsForCourse.length * 100);
//     }
//     course.ratingStats.teachingQuality = Object.keys(teachingQualityStats).map(key => ({
//         label: `${key}-star`, percentage: teachingQualityStats[Number(key)]
//     }));
//
//     course.ratingRubric.materialsQuality = ratingsForCourse.reduce((acc, curr) => acc + curr.ratingRubric.materialsQuality, 0) / ratingsForCourse.length;
//     course.ratingRubric.materialsQuality = Math.round(course.ratingRubric.materialsQuality * 10) / 10;
//     let materialsQualityStats: { [key: number]: number } = {1: 0, 2: 0, 3: 0, 4: 0, 5: 0};
//     ratingsForCourse.forEach(rating => {
//         materialsQualityStats[rating.ratingRubric.materialsQuality]++;
//     });
//     for (const rating in materialsQualityStats) {
//         materialsQualityStats[rating] = Math.round(materialsQualityStats[rating] / ratingsForCourse.length * 100);
//     }
//     course.ratingStats.materialsQuality = Object.keys(materialsQualityStats).map(key => ({
//         label: `${key}-star`, percentage: materialsQualityStats[Number(key)]
//     }));
//
//     course.ratingRubric.amountGains = ratingsForCourse.reduce((acc, curr) => acc + curr.ratingRubric.amountGains, 0) / ratingsForCourse.length;
//     course.ratingRubric.amountGains = Math.round(course.ratingRubric.amountGains * 10) / 10;
//     let amountGainsStats: { [key: number]: number } = {1: 0, 2: 0, 3: 0, 4: 0, 5: 0};
//     ratingsForCourse.forEach(rating => {
//         amountGainsStats[rating.ratingRubric.amountGains]++;
//     });
//     for (const rating in amountGainsStats) {
//         amountGainsStats[rating] = Math.round(amountGainsStats[rating] / ratingsForCourse.length * 100);
//     }
//     course.ratingStats.amountGains = Object.keys(amountGainsStats).map(key => ({
//         label: `${key}-star`, percentage: amountGainsStats[Number(key)]
//     }));
//
//     course.ratingRubric.difficulty = calcAvgTextScale(ratingsForCourse.map(rating => rating.ratingRubric.difficulty), difficultyScale);
//     let difficultyStats: { [key: string]: number } = {effortless: 0, manageable: 0, standard: 0, tough: 0, brutal: 0};
//     ratingsForCourse.forEach(rating => {
//         difficultyStats[rating.ratingRubric.difficulty]++;
//     });
//     for (const rating in difficultyStats) {
//         difficultyStats[rating] = Math.round(difficultyStats[rating] / ratingsForCourse.length * 100);
//     }
//     course.ratingStats.difficulty = Object.keys(difficultyStats).map(key => ({
//         label: key, percentage: difficultyStats[key]
//     }));
//
//     course.ratingRubric.workload = calcAvgTextScale(ratingsForCourse.map(rating => rating.ratingRubric.workload), workloadScale);
//     let workloadStats: { [key: string]: number } = {light: 0, moderate: 0, balanced: 0, heavy: 0, overwhelming: 0};
//     ratingsForCourse.forEach(rating => {
//         workloadStats[rating.ratingRubric.workload]++;
//     });
//     for (const rating in workloadStats) {
//         workloadStats[rating] = Math.round(workloadStats[rating] / ratingsForCourse.length * 100);
//     }
//     course.ratingStats.workload = Object.keys(workloadStats).map(key => ({
//         label: key, percentage: workloadStats[key]
//     }));
//
//     course.ratingRubric.grading = calcAvgTextScale(ratingsForCourse.map(rating => rating.ratingRubric.grading), gradingScale);
//     let gradingStats: { [key: string]: number } = {lenient: 0, flexible: 0, fair: 0, strict: 0, rigorous: 0};
//     ratingsForCourse.forEach(rating => {
//         gradingStats[rating.ratingRubric.grading]++;
//     });
//     for (const rating in gradingStats) {
//         gradingStats[rating] = Math.round(gradingStats[rating] / ratingsForCourse.length * 100);
//     }
//     course.ratingStats.grading = Object.keys(gradingStats).map(key => ({
//         label: key, percentage: gradingStats[key]
//     }));
// });
//
// // Calculate average rubric for each instructor
// instructors.forEach(instructor => {
//         const ratingsForInstructor: Rating[] = ratings.filter(rating => rating.courseOffering.taughtBy.includes(instructor));
//
//         instructor.ratingRubric.overallRating = ratingsForInstructor.reduce((acc, curr) => acc + curr.ratingRubric.overallRating, 0) / ratingsForInstructor.length;
//         instructor.ratingRubric.overallRating = Math.round(instructor.ratingRubric.overallRating * 10) / 10;
//
//         instructor.ratingRubric.teachingQuality = ratingsForInstructor.reduce((acc, curr) => acc + curr.ratingRubric.teachingQuality, 0) / ratingsForInstructor.length;
//         instructor.ratingRubric.teachingQuality = Math.round(instructor.ratingRubric.teachingQuality * 10) / 10;
//
//         instructor.ratingRubric.materialsQuality = ratingsForInstructor.reduce((acc, curr) => acc + curr.ratingRubric.materialsQuality, 0) / ratingsForInstructor.length;
//         instructor.ratingRubric.materialsQuality = Math.round(instructor.ratingRubric.materialsQuality * 10) / 10;
//
//         instructor.ratingRubric.amountGains = ratingsForInstructor.reduce((acc, curr) => acc + curr.ratingRubric.amountGains, 0) / ratingsForInstructor.length;
//         instructor.ratingRubric.amountGains = Math.round(instructor.ratingRubric.amountGains * 10) / 10;
//     }
// );
//
// // Post courses to localhost:4000/courses
// // courses.forEach(course => {
// //     auth("http://localhost:4000/courses", {
// //         method: "POST",
// //         headers: {
// //             "Content-Type": "application/json"
// //         },
// //         body: JSON.stringify({
// //             number: course.number,
// //             name: course.name,
// //             department: course.department,
// //             units: course.units,
// //             description: course.description,
// //         })
// //     })
// //         .then(response => response.json())
// //         .then(data => console.log(data))
// //         .catch(error => console.log(error));
// // })
//
// // Post instructors to localhost:4000/instructors
// // instructors.forEach(instructor => {
// //     auth("http://localhost:4000/instructors", {
// //         method: "POST",
// //         headers: {
// //             "Content-Type": "application/json"
// //         },
// //         body: JSON.stringify({
// //             name: instructor.name,
// //         })
// //     })
// //         .then(response => response.json())
// //         .then(data => console.log(data))
// //         .catch(error => console.log(error));
// // })
//
// // // Post users to localhost:4000/courseOfferings
// // users.forEach(user => {auth("http://localhost:4000/users", {
// //         method: "POST",
// //         headers: {
// //             "Content-Type": "application/json"
// //         },
// //         body: JSON.stringify({
// //             username: user.nickname,
// //             avatar_url: "https://static.vecteezy.com/system/resources/previews/020/902/060/non_2x/cute-sheep-cartoon-style-vector.jpg",
// //             andrew_id: faker.word.words(1),
// //             hashed_password: "123456",
// //         })
// //     })
// //         .then(response => response.json())
// //         .then(data => console.log(data))
// //         .catch(error => console.log(error));
// // })
//
// // // Post courseOfferings to localhost:4000/courseOfferings
// // courseOfferings.forEach(offering => {
// //     auth("http://localhost:4000/offerings", {
// //         method: "POST",
// //         headers: {
// //             "Content-Type": "application/json"
// //         },
// //         body: JSON.stringify({
// //             course_id: offering.course.id,
// //             semester: offering.semester,
// //             year: offering.year,
// //             location: offering.location,
// //         })
// //     })
// //         .then(response => response.json())
// //         .then(data => console.log(data))
// //         .catch(error => console.log(error));
// // })
//
// // // Post ratings to localhost:4000/ratings
// // ratings.forEach(rating => {
// //     auth("http://localhost:4000/ratings", {
// //         method: "POST",
// //         headers: {
// //             "Content-Type": "application/json"
// //         },
// //         body: JSON.stringify({
// //             user_id: rating.author.id,
// //             offering_id: rating.courseOffering.id,
// //             overall: rating.ratingRubric.overallRating,
// //             teaching: rating.ratingRubric.teachingQuality,
// //             materials: rating.ratingRubric.materialsQuality,
// //             gains: rating.ratingRubric.amountGains,
// //             difficulty: rating.ratingRubric.difficulty,
// //             workload: rating.ratingRubric.workload,
// //             grading: rating.ratingRubric.grading,
// //             comment: rating.comment,
// //         })
// //     })
// //         .then(response => response.json())
// //         .then(data => console.log(data))
// //         .catch(error => console.log(error));
// // })
//
// // // Post teaches to localhost:4000/teaches
// // instructors.forEach(instructor => {
// //     instructor.teaches.forEach(offering => {
// //         auth("http://localhost:4000/teaches", {
// //             method: "POST",
// //             headers: {
// //                 "Content-Type": "application/json"
// //             },
// //             body: JSON.stringify({
// //                 instructor_id: instructor.id,
// //                 offering_id: offering.id,
// //             })
// //         })
// //             .then(response => response.json())
// //             .then(data => console.log(data))
// //             .catch(error => console.log(error));
// //     })
// // })
