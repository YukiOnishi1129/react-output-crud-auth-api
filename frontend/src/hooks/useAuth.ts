import { useCallback, useState, useEffect } from "react";
import { useLocation, useNavigate } from "react-router";

import { NAVIGATION_PATH, NAVIGATION_LIST } from "../constants/navigation";
import {
  setAxiosAuthentication,
  removeAxiosAuthentication,
} from "../apis/globalAxios";
import { checkAuthentication } from "../apis/auth";
import { UserType } from "../types/User";

export const useAuth = () => {
  const navigate = useNavigate();
  const { pathname } = useLocation();
  const [user, setUser] = useState<UserType | null>(null);
  const [isAuth, setIsAuth] = useState<boolean>(false);

  const signIn = useCallback(
    (user: UserType, token: string) => {
      setUser(user);
      setIsAuth(true);
      setAxiosAuthentication(token);
      navigate(NAVIGATION_PATH.TOP);
    },
    [navigate]
  );

  const signOut = useCallback(() => {
    setUser(null);
    setIsAuth(false);
    removeAxiosAuthentication();
    navigate(NAVIGATION_PATH.LOGIN);
  }, [navigate]);

  const isExitBeforeAuthPage = useCallback(
    () =>
      pathname === NAVIGATION_PATH.SIGNUP || pathname === NAVIGATION_PATH.LOGIN,
    [pathname]
  );

  const authRouting = useCallback(async () => {
    let auth = false;
    const response = await checkAuthentication();
    if (response) {
      setUser(response.user);
      setIsAuth(true);
      auth = true;
    }

    // 未ログインでログイン後のページにいる場合、ログイン画面にリダイレクト
    if (!auth && !isExitBeforeAuthPage()) navigate(NAVIGATION_LIST.LOGIN);
    // ログイン済で未ログインのページにいる場合、Todo一覧ページにリダイレクト
    if (auth && isExitBeforeAuthPage()) navigate(NAVIGATION_LIST.TOP);
  }, [isExitBeforeAuthPage, navigate]);

  useEffect(() => {
    authRouting();
  }, [authRouting]);

  return {
    user,
    isAuth,
    signIn,
    signOut,
  };
};
