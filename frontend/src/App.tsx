import React, { useEffect, useState } from "react";
import * as runtime from '../wailsjs/runtime';
//import { PickCharacter } from "../wailsjs/go/main/App";

type Character = {
  CharId: number;
  CharStroke: number;
  JlptLvl: number;
  Char: string;
  ReadingJoyo: string;
  MeaningOn: string;
  MeaningKun: string;
};

const App: React.FC = () => {
    const [character, setCharacter] = useState<Character | null>(null);

    useEffect(() => {
        // Listening for character picked event
        const unlisten = runtime.EventsOn('characterPicked', (data: Character) => {
            setCharacter(data);
        });

        // Clean up on unmount
/*         return () => {
            unlisten();
        }; */
    }, []);

    if (!character) {
        return <div>Loading...</div>
    }

  return (
    <div
      style={{
        background: "#000",
        color: "#fff",
        width: "100vw",
        height: "100vh",
        position: "relative",
        display: "flex",
        flexDirection: "column",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      {/* Top right arrow button */}
      <button
        style={{
          position: "absolute",
          top: 24,
          right: 24,
          background: "none",
          border: "none",
          color: "#fff",
          fontSize: "2rem",
          cursor: "pointer",
        }}
        onClick={() => {}}
        aria-label="Next Character"
      >
        ↗
      </button>

      {/* Centered Kanji character */}
      <div
        style={{
          fontSize: "10rem",
          fontWeight: "bold",
          textAlign: "center",
          margin: "auto",
        }}
      >
        {character?.Char || ""}
      </div>

      {/* Bottom stacked texts */}
      <div
        style={{
          position: "absolute",
          bottom: 48,
          left: 0,
          width: "100%",
          textAlign: "center",
        }}
      >
        <div style={{ fontSize: "1.5rem", marginBottom: 8 }}>
          Reading: {character?.ReadingJoyo || ""}
        </div>
        <div style={{ fontSize: "1.2rem", opacity: 0.8 }}>
          Meaning: {character?.MeaningOn || ""}
        </div>
      </div>
    </div>
  );
};

export default App;
