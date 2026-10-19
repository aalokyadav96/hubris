export const coreRoutes = [
  {
    path: "/",
    component: () => import("../../pages/home.js"),
    functionName: "Home",
    meta: {
      title: "Home"
    }
  },

  {
    path: "/home",
    component: () => import("../../pages/home.js"),
    functionName: "Home",
    meta: {
      title: "Home"
    }
  },

  {
    path: "/login",
    component: () => import("../../pages/auth/auth.js"),
    functionName: "Auth",
    meta: {
      guestOnly: true,
      title: "Login"
    }
  },

  {
    path: "/profile",
    component: () => import("../../pages/profile/userProfile.js"),
    functionName: "MyProfile",
    meta: {
      requiresAuth: true,
      title: "My Profile"
    }
  },

  {
    path: "/user/:id",
    component: () => import("../../pages/profile/userProfile.js"),
    functionName: "UserProfile",
    meta: {
      title: "User Profile"
    }
  },

  {
    path: "/settings",
    component: () => import("../../pages/profile/settings.js"),
    functionName: "Settings",
    meta: {
      requiresAuth: true,
      title: "Account Settings"
    }
  },

  {
    path: "/map",
    component: () => import("../../pages/gtamap/mapgta.js"),
    functionName: "MapGTA",
    meta: {
      title: "Map"
    }
  },

  {
    path: "/cart",
    component: () => import("../../pages/cart/cart.js"),
    functionName: "Cart",
    meta: {
      requiresAuth: true,
      title: "Shopping Cart"
    }
  },

  {
    path: "/my-orders",
    component: () => import("../../pages/cart/myorders.js"),
    functionName: "MyOrders",
    meta: {
      requiresAuth: true,
      title: "My Orders"
    }
  }
];
