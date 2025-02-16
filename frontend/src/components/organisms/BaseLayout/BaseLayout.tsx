import { FC, ReactNode } from "react";

import { Navigation } from "../../molecules/Navigation";
import styles from "./style.module.css";

type BaseLayoutProps = {
  children: ReactNode;
  title: string;
};

export const BaseLayout: FC<BaseLayoutProps> = ({ children, title }) => (
  <div>
    <header className={styles.header}>
      <Navigation />
    </header>
    <main className={styles.container}>
      <h1 className={styles.title}>{title}</h1>
      {children}
    </main>
  </div>
);
