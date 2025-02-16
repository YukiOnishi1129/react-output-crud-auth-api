import { useCallback, useEffect } from "react";
import { useParams, useNavigate } from "react-router";
import { useForm } from "react-hook-form";
import * as z from "zod";
import { zodResolver } from "@hookform/resolvers/zod";

import { getTodo, updateTodo } from "../../../apis/todo";
import { NAVIGATION_PATH } from "../../../constants/navigation";

const schema = z.object({
  title: z
    .string()
    .min(1, "タイトルは必須です。")
    .max(10, "10文字以内で入力してください。"),
  content: z.string().optional(),
});

export const useTodoEditTemplate = () => {
  const { id } = useParams();
  const navigate = useNavigate();

  const {
    control,
    handleSubmit,
    formState: { errors },
    setValue,
  } = useForm<z.infer<typeof schema>>({
    resolver: zodResolver(schema),
  });

  const fetchTodo = useCallback(async () => {
    if (!id) return;
    const response = await getTodo({ id });
    if (!response.data) return;
    setValue("title", response.data.title);
    setValue("content", response.data.content);
  }, [id, setValue]);

  const handleEditSubmit = handleSubmit(
    useCallback(
      async (values: z.infer<typeof schema>) => {
        if (!id) return;
        await updateTodo({
          id,
          title: values.title,
          content: values.content,
        });
        navigate(NAVIGATION_PATH.TOP);
      },
      [navigate, id]
    )
  );

  useEffect(() => {
    fetchTodo();
  }, [fetchTodo]);

  return {
    control,
    errors,
    handleEditSubmit,
  };
};
