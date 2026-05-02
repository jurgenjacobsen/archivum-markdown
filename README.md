![Archivum Markdown - Banner](/frontend/public/banner.png)

[![Wails build](https://github.com/jurgenjacobsen/archivum-markdown/actions/workflows/build.yml/badge.svg)](https://github.com/jurgenjacobsen/archivum-markdown/actions/workflows/build.yml)
[![wakatime](https://wakatime.com/badge/user/010adc07-6382-419f-87bc-0b3f507ee495/project/fb303ffe-8258-4dcc-bfc9-e34da048aafb.svg)](https://wakatime.com/badge/user/010adc07-6382-419f-87bc-0b3f507ee495/project/fb303ffe-8258-4dcc-bfc9-e34da048aafb)
![GitHub last commit (branch)](https://img.shields.io/github/last-commit/jurgenjacobsen/archivum-markdown/main)
![GitHub top language](https://img.shields.io/github/languages/top/jurgenjacobsen/archivum-markdown)
![GitHub repo size](https://img.shields.io/github/repo-size/jurgenjacobsen/archivum-markdown)
![GitHub Downloads (all assets, all releases)](https://img.shields.io/github/downloads/jurgenjacobsen/archivum-markdown/total)

# Archivum Markdown
**Archivum Markdown** is a high-performance desktop Markdown editor designed for speed, simplicity, and focus. Built with Wails, it combines the power of Go with the flexibility of React and TypeScript to provide a seamless writing experience.

## Key Features
- **Workspace Integration:** Open any folder as a workspace and navigate your files with an integrated sidebar.
- **Live Preview:** Real-time Markdown rendering as you type.
- **Synchronized Scrolling:** Keep your editor and preview perfectly in sync.
- **Rich Formatting:** Quick access to common Markdown syntax (Headers, Lists, Checklists, Code Blocks, etc.).
- **Auto-Save:** Focus on writing while Archivum ensures your progress is safely stored.
- **File Management:** Create, rename, and delete files and directories directly within the app.
- **Print Support:** Easily export your documents to PDF or print them.
- **Fast & Lightweight:** Native performance with a modern, distraction-free interface.

## Tech Stack
- **Backend:** [Go](https://go.dev/) with [Wails](https://wails.io/)
- **Frontend:** [React](https://reactjs.org/), [TypeScript](https://www.typescriptlang.org/), [Tailwind CSS](https://tailwindcss.com/)
- **State Management:** React Hooks
- **Icons & Assets:** Custom branding by Archivum

## Getting Started
### Prerequisites

- [Go](https://go.dev/dl/) (v1.23+)
- [Node.js](https://nodejs.org/) & [NPM](https://www.npmjs.com/)
- [Wails CLI](https://wails.io/docs/gettingstarted/installation)

### Live Development
To run in live development mode, run the following command in the project directory:

```bash
wails dev
```

This will start a Vite development server for the frontend and recompile the Go backend on changes.

### Building
To build a redistributable, production-ready package:

```bash
wails build
```

### Delivery
```
git tag v1.0.1
git push origin v1.0.1
```

The resulting binary will be located in the `build/bin` directory.

## Project Structure

- `app.go`: Main backend logic and Go-to-Frontend bindings.
- `main.go`: Application entry point and configuration.
- `frontend/`: React-based user interface.
- `wails.json`: Project configuration and metadata.

## License

Copyright (c) 2026 Archivum. All rights reserved.
