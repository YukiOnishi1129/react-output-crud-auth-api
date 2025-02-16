import { FC } from "react";
import { Controller } from "react-hook-form";

import { InputFormSection } from "../../molecules";
import { CommonButton } from "../../atoms";

import { useSignUpTemplate } from "./useSignUpTemplate";

import styles from "./style.module.css";

export const SignUpTemplate: FC = () => {
  const { control, errors, handleRegisterSubmit } = useSignUpTemplate();

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>SignUp</h1>
      <form className={styles.form} onSubmit={handleRegisterSubmit}>
        <div className={styles.area}>
          <Controller
            name="name"
            render={({ field }) => (
              <InputFormSection
                type="text"
                placeholder="user name"
                errorMessage={errors.name?.message}
                {...field}
              />
            )}
            control={control}
          />
        </div>

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
          <CommonButton type="submit">{"Sign Up"}</CommonButton>
        </div>
      </form>
    </div>
  );
};
