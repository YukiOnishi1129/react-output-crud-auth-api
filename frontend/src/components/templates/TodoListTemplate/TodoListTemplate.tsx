import { Controller } from "react-hook-form";
import { InputFormSection } from "../../molecules";
import { BaseLayout, TodoList } from "../../organisms";

import { useTodoListTemplate } from "./useTodoListTemplate";

import styles from "./style.module.css";

export const TodoListTemplate = () => {
  const { control, showTodoList, handleDeleteTodo } = useTodoListTemplate();

  return (
    <BaseLayout title={"TodoList"}>
      <div className={styles.container}>
        {/* Todo検索フォームエリア */}
        <div className={styles.area}>
          <Controller
            name="keyword"
            render={({ field }) => (
              <InputFormSection placeholder={"Search Keyword"} {...field} />
            )}
            control={control}
          />
        </div>
        {/* Todoリスト一覧表示 */}
        <div className={styles.area}>
          {showTodoList.length > 0 && (
            <TodoList todoList={showTodoList} onDeleteTodo={handleDeleteTodo} />
          )}
        </div>
      </div>
    </BaseLayout>
  );
};
