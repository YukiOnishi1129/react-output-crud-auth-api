/**
 * ベースPATH
 */
export const BASE_PATH = "";

/**
 * リンク先一覧
 * 遷移先定義の際に使用
 */
export const NAVIGATION_LIST = {
  LOGIN: `${BASE_PATH}`,
  SIGNUP: `${BASE_PATH}/signup`,
  TOP: `${BASE_PATH}/todo`,
  DETAIL: `${BASE_PATH}/todo/detail/:id`,
  CREATE: `${BASE_PATH}/todo/create`,
  EDIT: `${BASE_PATH}/todo/edit/:id`,
};

/**
 * パス一覧
 * 画面遷移時の使用
 */
export const NAVIGATION_PATH = {
  LOGIN: `${BASE_PATH}`,
  SIGNUP: `${BASE_PATH}/signup`,
  TOP: `${BASE_PATH}/todo`,
  DETAIL: `${BASE_PATH}/todo/detail/`,
  CREATE: `${BASE_PATH}/todo/create`,
  EDIT: `${BASE_PATH}/todo/edit/`,
};
