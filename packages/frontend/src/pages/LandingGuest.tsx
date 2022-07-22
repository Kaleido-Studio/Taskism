import { Component } from "solid-js";
import './index.scss';

export const LandingGuest:Component = () => {
    return (<>
    <div class="container border grid grid-cols-2 <sm:grid-cols-1 mx-auto relative top-34">
        <div class="title border h-30 col-span-1">
            <h1 class="text-5xl antialiased font-extrabold text-center title-color">Taskism</h1>
        </div>
        <div class="form border h-60 col-span-1 row-span-2">
            <form>
                
            </form>
        </div>
        <div class="desc border h-30 col-span-1 row-start-2 font-semibold text-center">
            Unleash your potential and make your life easier.
        </div>
    </div>
    </>)
}