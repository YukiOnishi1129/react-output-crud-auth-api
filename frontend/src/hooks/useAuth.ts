import { useCallback, useState } from "react";
import { useLocation } from "react-router";

import { NAVIGATION_PATH } from "../constants/navigation";
import { UserType } from "../types/User";

export const useAuth = () => {
  const { pathname } = useLocation();
  const [user, setUser] = useState<UserType | null>(null);
  const [isAuth, setIsAuth] = useState<boolean>(false);

  const signIn = useCallback((user: UserType) => {
    setUser(user);
    setIsAuth(true);
  }, []);

  const signOut = useCallback(() => {
    setUser(null);
    setIsAuth(false);
  }, []);

  const isExitBeforeAuthPage = useCallback(
    () =>
      pathname === NAVIGATION_PATH.SIGNUP || pathname === NAVIGATION_PATH.LOGIN,
    [pathname]
  );

  const authRouting = useCallback(() => {
    let auth = false;
  }, []);

  return {
    user,
    isAuth,
    signIn,
    signOut,
  };
};
