import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    redirect: { name: "Dashboard" },
  },
  {
    path: "/dashboard",
    name: "Dashboard",
    component: () => import("../views/Dashboard.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/namespaces",
    name: "Namespaces",
    component: () => import("../views/Namespaces.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/accounts",
    name: "Accounts",
    component: () => import("../views/Accounts.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/sp",
    name: "ServicesProviders",
    component: () => import("../views/ServicesProviders.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/sp/create",
    name: "ServicesProviders create",
    component: () => import("../views/ServicesProvidersCreate.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/sp/:uuid",
    name: "ServicesProvider",
    component: () => import("../views/ServicesProvidersPage.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/dns",
    name: "DNS manager",
    component: () => import("../views/dnsManager.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/dns/:dnsname",
    name: "Zone manager",
    component: () => import("../views/ZoneManager.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/settigns",
    name: "Settings",
    component: () => import("../views/Settings.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/services",
    name: "Services",
    component: () => import("../views/Services.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/services/create",
    name: "Service create",
    component: () => import("../views/ServiceCreate.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/services/:serviceId",
    name: "Service",
    component: () => import("../views/ServicePage.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/login",
    name: "Login",
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/login.vue"),
    meta: {
      requireUnlogin: true,
    },
  },
];

const router = new VueRouter({
  // mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
