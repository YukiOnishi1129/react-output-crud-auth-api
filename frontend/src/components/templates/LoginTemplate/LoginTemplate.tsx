import { FC } from "react";
import { NavLink } from "react-router";
import { Controller } from "react-hook-form";

import { NAVIGATION_LIST } from "../../../constants/navigation";

import { InputFormSection } from "../../molecules";
import { CommonButton } from "../../atoms";

import { useLoginTemplate } from "./useLoginTemplate";

import styles from "./style.module.css";

export const LoginTemplate: FC = () => {
  const { control, errors, handleLoginSubmit } = useLoginTemplate();

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Login</h1>
      <form className={styles.form} onSubmit={handleLoginSubmit}>
        <div className={styles.area}>
          <Controller
            name="email"
            render={({ field }) => (
              <InputFormSection
                type="email"
                placeholder="email"
                errorMessage={errors.email?.message}
                {...field}
              />
            )}
            control={control}
          />
        </div>
        <div className={styles.area}>
          <Controller
            name="password"
            render={({ field }) => (
              <InputFormSection
                type="password"
                placeholder="password"
                errorMessage={errors.password?.message}
                {...field}
              />
            )}
            control={control}
          />
        </div>
        <div className={styles.area}>
          <Controller
            name="password_confirmation"
            render={({ field }) => (
              <InputFormSection
                type="password"
                placeholder="confirm password"
                errorMessage={errors.password_confirmation?.message}
                {...field}
              />
            )}
            control={control}
          />
        </div>
        <div className={styles.area}>
          <CommonButton type="submit">{"Login"}</CommonButton>
        </div>
        <div className={styles.link}>
          <NavLink to={NAVIGATION_LIST.SIGNUP}>&lt;&lt; to signup page</NavLink>
        </div>
      </form>
    </div>
  );
};
