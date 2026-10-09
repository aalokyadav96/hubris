export const legalRoutes = [
  {
    path: "/about",
    component: () => import("../../pages/legalPages/home.js"),
    functionName: "About",
    meta: {
      title: "About Us"
    }
  },

  {
    path: "/contact",
    component: () => import("../../pages/legalPages/home.js"),
    functionName: "Contact",
    meta: {
      title: "Contact Us"
    }
  },

  {
    path: "/faq",
    component: () => import("../../pages/legalPages/home.js"),
    functionName: "Faq",
    meta: {
      title: "FAQ"
    }
  },

  {
    path: "/terms",
    component: () => import("../../pages/legalPages/home.js"),
    functionName: "Terms",
    meta: {
      title: "Terms of Service"
    }
  },

  {
    path: "/privacy",
    component: () => import("../../pages/legalPages/home.js"),
    functionName: "Privacy",
    meta: {
      title: "Privacy Policy"
    }
  },

  {
    path: "/refund",
    component: () => import("../../pages/legalPages/home.js"),
    functionName: "Refund",
    meta: {
      title: "Refund Policy"
    }
  },

  {
    path: "/shipping",
    component: () => import("../../pages/legalPages/home.js"),
    functionName: "Shipping",
    meta: {
      title: "Shipping Information"
    }
  },

  {
    path: "/returns",
    component: () => import("../../pages/legalPages/home.js"),
    functionName: "Returns",
    meta: {
      title: "Returns Policy"
    }
  },

  {
    path: "/disclaimer",
    component: () => import("../../pages/legalPages/home.js"),
    functionName: "Disclaimer",
    meta: {
      title: "Disclaimer"
    }
  },

  {
    path: "/blog",
    component: () => import("../../pages/legalPages/home.js"),
    functionName: "Blog",
    meta: {
      title: "Blog"
    }
  }
];
