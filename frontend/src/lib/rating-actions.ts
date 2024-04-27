// import type {Rating} from "$lib/types";
// import {goto} from "$app/navigation";
//
// export const onLike = (rating: Rating) => {
//     if (rating.likedByUser) {
//         rating.likes--;
//     } else {
//         rating.likes++;
//     }
//     rating.likedByUser = !rating.likedByUser;
// }
//
// export const onDislike = (rating: Rating) => {
//     if (rating.dislikedByUser) {
//         rating.dislikes--;
//     } else {
//         rating.dislikes++;
//     }
//     rating.dislikedByUser = !rating.dislikedByUser;
// }
//
// export const onReadMore = (url: string, ratingID: string) => {
//     goto(url)
//         .then(() => {
//             let ratingElement = document.getElementById(ratingID);
//             if (ratingElement) {
//                 window.scrollTo({top: ratingElement.offsetTop, behavior: "smooth"});
//             }
//         })
//         .catch((error) => console.error(error));
// }