<script lang="ts">
    import '../app.pcss';
    import Header from "$lib/header/Header.svelte";
    import ScrollToTop from "$lib/scroll-to-top/ScrollToTop.svelte";
    import Footer from "$lib/footer/Footer.svelte";
    import Toasts from "$lib/toast/Toasts.svelte";
    import LoginModal from "$lib/auth/LoginModal.svelte";
    import type {User} from "$lib/types";

    export let data : {token: string | null, expiry: string | null, user : User | null};
</script>

<Header token={data.token} expiry={data.expiry} user={data.user}/>

<div class="relative container flex flex-col mx-auto px-4 text-gray-900 min-h-[calc(100vh-274px)] lg:min-h-[calc(100vh-218px)]">
    <slot/>
</div>

<Footer/>

<ScrollToTop/>
<LoginModal/>
<Toasts/>
