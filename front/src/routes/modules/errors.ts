export const errorRoutes = [
  {
    path: "/404",
    component: () => import("../../pages/errors/error.js"),
    functionName: "NotFound",
    meta: {
      title: "Page Not Found"
    }
  },

  {
    path: "/error/404",
    component: () => import("../../pages/errors/error.js"),
    functionName: "NotFound",
    meta: {
      title: "Page Not Found"
    }
  },

  {
    path: "/403",
    component: () => import("../../pages/errors/error.js"),
    functionName: "Forbidden",
    meta: {
      title: "Access Denied"
    }
  },

  {
    path: "/error/403",
    component: () => import("../../pages/errors/error.js"),
    functionName: "Forbidden",
    meta: {
      title: "Access Denied"
    }
  }
];
