import { FC, useCallback } from "react";
import { NavLink, useNavigate } from "react-router";
import { Controller, useForm } from "react-hook-form";
import * as z from "zod";
import { zodResolver } from "@hookform/resolvers/zod";

import { setAxiosAuthentication } from "../../../apis/globalAxios";
import { login } from "../../../apis/auth";

import {
  NAVIGATION_LIST,
  NAVIGATION_PATH,
} from "../../../constants/navigation";

import { InputFormSection } from "../../molecules";
import { CommonButton } from "../../atoms";

import styles from "./style.module.css";

const schema = z.object({
  email: z.string().email("メールアドレスの形式で入力してください"),
  password: z.string().min(8, "8文字以上で入力してください"),
  password_confirmation: z.string().min(8, "8文字以上で入力してください"),
});

export const LoginTemplate: FC = () => {
  const navigate = useNavigate();

  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm<z.infer<typeof schema>>({
    resolver: zodResolver(schema),
  });

  const handleLoginSubmit = handleSubmit(
    useCallback(
      async (values: z.infer<typeof schema>) => {
        const { email, password } = values;
        const res = await login(email, password);
        if (res) {
          setAxiosAuthentication(res.token);
          navigate(NAVIGATION_PATH.TOP);
        }
      },
      [navigate]
    )
  );

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
                placeholder="password"
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
