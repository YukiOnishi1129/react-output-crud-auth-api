import { Controller } from "react-hook-form";

import { BaseLayout } from "../../organisms";
import { InputFormSection, TextAreaSection } from "../../molecules";
import { CommonButton } from "../../atoms";

import { useTodoCreateTemplate } from "./useTodoCreateTemplate";

import styles from "./style.module.css";

export const TodoCreateTemplate = () => {
  const { control, errors, handleAddSubmit } = useTodoCreateTemplate();

  return (
    <BaseLayout title={"Create Todo"}>
      <div></div>
      <form className={styles.container} onSubmit={handleAddSubmit}>
        <div className={styles.area}>
          <Controller
            name="title"
            render={({ field }) => (
              <InputFormSection
                placeholder={"Title"}
                errorMessage={errors.title?.message}
                {...field}
              />
            )}
            control={control}
          />
        </div>
        <div className={styles.area}>
          <Controller
            name="content"
            render={({ field }) => (
              <TextAreaSection
                placeholder={"Content"}
                errorMessage={errors.content?.message}
                {...field}
              />
            )}
            control={control}
          />
        </div>
        <div className={styles.area}>
          <CommonButton type="submit">{"Create Todo"}</CommonButton>
        </div>
      </form>
    </BaseLayout>
  );
};
