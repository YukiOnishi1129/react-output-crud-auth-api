import { Controller } from "react-hook-form";

import { BaseLayout } from "../../organisms";
import { InputFormSection, TextAreaSection } from "../../molecules";
import { CommonButton } from "../../atoms";
import { useTodoEditTemplate } from "./useTodoEditTemplate";

import styles from "./style.module.css";

export const TodoEditTemplate = () => {
  const { control, errors, handleEditSubmit } = useTodoEditTemplate();

  return (
    <BaseLayout title={"TodoEdit"}>
      <form className={styles.container} onSubmit={handleEditSubmit}>
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
          <CommonButton type="submit">{"Edit Todo"}</CommonButton>
        </div>
      </form>
    </BaseLayout>
  );
};
