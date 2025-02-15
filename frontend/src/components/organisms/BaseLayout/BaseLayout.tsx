import { FC, ReactNode } from "react";

import { Navigation } from "../../molecules/Navigation";
import styles from "./style.module.css";

type BaseLayoutProps = {
  children: ReactNode;
  title: string;
};

export const BaseLayout: FC<BaseLayoutProps> = ({ children, title }) => (
  <div className={styles.container}>
    <section className={styles.common}>
      <Navigation />
    </section>
    <h1 className={styles.title}>{title}</h1>
    {children}
  </div>
);
