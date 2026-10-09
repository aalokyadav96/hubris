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
  },

  {
    path: "/deliveries",
    component: () => import("../../pages/delivery/deliveries.js"),
    functionName: "Deliveries",
    meta: {
      requiresAuth: true,
      title: "Deliveries"
    }
  },

  {
    path: "/delivery/create",
    component: () => import("../../pages/delivery/createDelivery.js"),
    functionName: "Createdelivery",
    meta: {
      requiresAuth: true,
      title: "Create Delivery"
    }
  },

  {
    path: "/delivery/track/:id",
    component: () => import("../../pages/delivery/trackDelivery.js"),
    functionName: "TrackDelivery",
    meta: {
      requiresAuth: true,
      title: "Track Delivery"
    }
  },

  {
    path: "/delivery/:id",
    component: () => import("../../pages/delivery/displayDelivery.js"),
    functionName: "Delivery",
    meta: {
      requiresAuth: true,
      title: "Delivery"
    }
  },

  {
    path: "/dash/driver",
    component: () => import("../../pages/delivery/driverDash.js"),
    functionName: "DriverDash",
    meta: {
      requiresAuth: true,
      roles: ["driver", "admin"],
      roleMatchMode: "ANY",
      title: "Driver Dashboard"
    }
  },

  {
    path: "/wallet",
    component: () => import("../../pages/wallet/wallet.js"),
    functionName: "Wallet",
    meta: {
      requiresAuth: true,
      title: "Wallet"
    }
  }
];
