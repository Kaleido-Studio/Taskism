import { Component } from "solid-js";
import './index.scss';

export const LandingGuest:Component = () => {
    return (<>
<div class="hero container min-h-screen mx-auto">
  <div class="hero-content flex-col lg:flex-row container justify-around w-screen">
    <div class="text-center lg:text-left">
      <h1 class="text-6xl font-bold title-color">Taskism</h1>
      <p class="py-7 font-semibold">Unleash your potential</p>
    </div>
    <div class="card flex-shrink-0 w-full max-w-sm shadow-2xl bg-base-100">
      <div class="card-body">
        <div class="form-control">
          <label class="label">
            <span class="label-text">Email</span>
          </label>
          <input type="text" placeholder="email" class="input input-bordered" />
        </div>
        <div class="form-control">
          <label class="label">
            <span class="label-text">Password</span>
          </label>
          <input type="text" placeholder="password" class="input input-bordered" />
          <label class="label">
            <a href="#" class="label-text-alt link link-hover">Forgot password?</a>
          </label>
        </div>
        <div class="form-control mt-6">
          <button class="btn btn-primary">Login</button>
        </div>
      </div>
    </div>
  </div>
</div>
    </>)
}