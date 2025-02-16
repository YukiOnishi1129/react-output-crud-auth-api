import { NavigationLink } from "../../atoms/NavigationLink";
import { NAVIGATION_PATH } from "../../../constants/navigation";
import { useAuthContext } from "../../../hooks/useAuthContext";
import styles from "./style.module.css";

export const Navigation = () => {
  const { signOut } = useAuthContext();
  return (
    <div className={styles.header}>
      <h1 className={styles.title}>Todo List</h1>
      <nav className={styles.nav}>
        <ul className={styles.ul}>
          <NavigationLink title={"Top"} linkPath={NAVIGATION_PATH.TOP} />
          <NavigationLink title={"Create"} linkPath={NAVIGATION_PATH.CREATE} />
          <li className={styles.li}>
            <button className={styles.button} onClick={() => signOut()}>
              SignOut
            </button>
          </li>
        </ul>
      </nav>
    </div>
  );
};
