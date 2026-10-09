import { getCurrentAllowedFeatures } from "../../config/domainFeatures.js";
import { metaGuard } from "../../middleware/middleware.js";
import { adminRoutes } from "./admin.js";
import { farmsRoutes } from "./farms.js";
import { eventsRoutes } from "./events.js";
import { baitoRoutes } from "./baito.js";
import { socialRoutes } from "./social.js";
import { chatsRoutes } from "./chats.js";
import { placesRoutes } from "./places.js";
import { coreRoutes } from "./core.js";
import { legalRoutes } from "./legal.js";
import { errorRoutes } from "./errors.js";

interface RegistryRoute {
  path: string;
  component: () => Promise<any>;
  functionName?: string;
  meta?: Record<string, any>;
  middleware?: Array<(context: any) => Promise<boolean | string | void> | boolean | string | void>;
  beforeEnter?: (context: any) => Promise<boolean | string | void> | boolean | string | void;
  afterEnter?: (context: any) => Promise<void> | void;
}

const featureModules: Record<string, RegistryRoute[]> = {
  admin: adminRoutes,
  farms: farmsRoutes,
  events: eventsRoutes,
  baito: baitoRoutes,
  social: socialRoutes,
  chats: chatsRoutes,
  places: placesRoutes
};

function routeSpecificity(route: RegistryRoute): number {
  const path = String(route.path || "");
  const segments = path.split("/").filter(Boolean);

  let score = 0;

  for (const segment of segments) {
    if (segment.startsWith("*")) {
      score += 1;
    } else if (segment.startsWith(":")) {
      score += 10;
    } else {
      score += 100;
    }
  }

  score += segments.length;
  return score;
}

export function buildRoutes(): RegistryRoute[] {
  const allowedFeatures: string[] = getCurrentAllowedFeatures();

  const aggregatedRoutes: RegistryRoute[] = [
    ...coreRoutes,
    ...legalRoutes,
    ...errorRoutes
  ];

  Object.entries(featureModules).forEach(([featureKey, routesList]) => {
    const enabled =
      allowedFeatures.includes("ALL") ||
      allowedFeatures.includes(featureKey);

    if (enabled && Array.isArray(routesList)) {
      aggregatedRoutes.push(...routesList);
    }
  });

  if (allowedFeatures.includes("ALL") || allowedFeatures.includes("admin")) {
    aggregatedRoutes.push({
      path: "/admin/*path",
      component: () => import("../../pages/admin/dashboard.js"),
      functionName: "AdminDashboard",
      meta: {
        requiresAuth: true,
        roles: ["admin"],
        title: "Admin Portal"
      }
    });
  }

  return aggregatedRoutes
    .map((route) => ({
      ...route,
      middleware: [metaGuard, ...(route.middleware || [])]
    }))
    .sort((a, b) => routeSpecificity(b) - routeSpecificity(a));
}
