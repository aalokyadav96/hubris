import {
  metaGuard,
  authGuard,
  guestGuard,
  roleGuard,
  permissionGuard,
  onboardingGuard,
  featureFlagGuard,
  unsavedChangesGuard,
  titleGuard,
  analyticsGuard
} from "../middleware/middleware.js";
import { buildRoutes } from "./modules/registry.js";

/* =========================================================
   TYPES & INTERFACES
========================================================= */

export type MiddlewareFn = (context: any) => Promise<boolean | string | void> | boolean | string | void;

export interface RouteMeta {
  title?: string;
  requiresAuth?: boolean;
  guestOnly?: boolean;
  roles?: string[];
  roleMatchMode?: "ANY" | "ALL";
  permissions?: string[];
  featureFlag?: string;
  [key: string]: any;
}

export interface Route {
  path: string;
  component: () => Promise<any>;
  functionName?: string;
  meta?: RouteMeta;
  middleware?: MiddlewareFn[];
  beforeEnter?: (context: any) => Promise<boolean | string | void> | boolean | string | void;
  afterEnter?: (context: any) => Promise<void> | void;
}


export const routes: Route[] = buildRoutes();

// Export granular middleware helpers for standalone usage if needed
export {
  metaGuard,
  authGuard,
  guestGuard,
  roleGuard,
  permissionGuard,
  onboardingGuard,
  featureFlagGuard,
  unsavedChangesGuard,
  titleGuard,
  analyticsGuard
};