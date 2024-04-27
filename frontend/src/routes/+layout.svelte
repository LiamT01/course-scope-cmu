<script lang="ts">
    import '../app.pcss';
    import Header from "$lib/header/Header.svelte";
    import ScrollToTop from "$lib/scroll-to-top/ScrollToTop.svelte";
    import Footer from "$lib/footer/Footer.svelte";
    import Toasts from "$lib/toast/Toasts.svelte";
    import LoginModal from "$lib/auth/LoginModal.svelte";
</script>

<Header/>

<div class="relative container flex flex-col mx-auto px-4 text-gray-900 min-h-[calc(100vh-274px)] lg:min-h-[calc(100vh-218px)]">
    <slot/>
</div>

<Footer/>

<ScrollToTop/>
<LoginModal/>
<Toasts/>
