import { FC } from "react";
import { NavLink } from "react-router";
import styles from "./style.module.css";

type NavigationLinkProps = {
  title: string;
  linkPath: string;
};

export const NavigationLink: FC<NavigationLinkProps> = ({
  title,
  linkPath,
}) => (
  <li className={styles.li}>
    <NavLink to={linkPath}>{title}</NavLink>
  </li>
);
