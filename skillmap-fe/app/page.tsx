import Image from "next/image";
import styles from "./page.module.css";

export default function Home() {
  return (
    <main className={styles.main}>
      <div className={styles.description}>
        <p>
          <code className={styles.code}><a href="https://youtu.be/pHtohACmyj8">YouTube.com/@buildfromzero</a></code>
        </p>
        <div>
          <a
            href="https://vercel.com?utm_source=create-next-app&utm_medium=appdir-template&utm_campaign=create-next-app"
            target="_blank"
            rel="noopener noreferrer"
          >
            By Build From Zero
          </a>
        </div>
      </div>

      <div className={styles.center}>
        <Image
          className={styles.logo}
          src="/gin.png"
          alt="Build FRom zero Logo"
          width={600}
          height={300}
          priority
        />
      </div>

      <div className={styles.grid}>
        <a
          href=""
          className={styles.card}
          rel="noopener noreferrer"
        >
          <h2>
            Users <span>-&gt;</span>
          </h2>
          <p>Manage Users</p>
        </a>

        <a
          href=""
          className={styles.card}
          rel="noopener noreferrer"
        >
          <h2>
            Search <span>-&gt;</span>
          </h2>
          <p>Search and filter</p>
        </a>

        <a
          href=""
          className={styles.card}
          rel="noopener noreferrer"
        >
          <h2>
            Skills <span>-&gt;</span>
          </h2>
          <p>Manage Skills</p>
        </a>

        <a
          href=""
          className={styles.card}
          rel="noopener noreferrer"
        >
          <h2>
            Settings <span>-&gt;</span>
          </h2>
          <p>
            Manage App settings
          </p>
        </a>
      </div>
    </main>
  );
}
