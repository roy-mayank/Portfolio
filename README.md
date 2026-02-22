# Portfolio Application - SvelteKit

Another revamped version of a portfolio website, currently in progress. This version inherits a lot of the same features of 'Movie Game' TMDB API calls based ratings game, firebase hosted blog. Changes include the use of a newstack (SvelteKit), newer design and up to date information.
Work in progress include the addition of Go based functionality for real time user information tracking and for a collaborative document that members can work on simultaneously. I also intend to proactively display hobbies engaged in, and to churn out more written content.

Find the [old repo here](https://github.com/roy-mayank/new-portfolio-app)

## Deployment

The project has been deployed [here](https://roys-latest-portfolio.vercel.app) using Vercel. Please do report to [my email](roy050703@gmail.com) in case of any bugs or errors (I would sincerely appreciate that).

## Features expected when complete
- Legacy TMDB API based game
- Authentication options for commenting and/or liking posted content
- Blog page with technical and non-technical writings
- Shared collaboratory document space (to be refined further)

## Getting started locally

To run the project locally:
- run ```pnpm run dev``` on downloaded repo's /frontend directory
- Then use the [scripts](#Scripts)

### Go view-counter (optional)

A small Go service records per-page view counts. The frontend can display "This page: N views" when the API is configured.

**Run Go locally:**
- From the repo root: `cd go && go run .` (listens on `http://localhost:8080`)
- In the frontend directory, set env `PUBLIC_GO_API_URL=http://localhost:8080` (e.g. in a `.env` file or when running `pnpm run dev`). Then the SvelteKit app will ping the Go API on each page load and show the count at the bottom.

**Endpoints:**
- `GET /api/view?path=/some/path` – increment count for that path, returns `{"path":"/some/path","count":N}`
- `GET /api/stats` – returns all path counts as JSON (e.g. to inspect traffic)

**Deploy Go:** Deploy the `go/` app to Fly.io, Railway, or Render. Set `PUBLIC_GO_API_URL` in your frontend host (Vercel/Netlify) to that URL so the live site sends view pings to your Go service. Counts are in-memory by default (resets on restart); you can add file persistence later if needed.

## Author

Mayank Roy - CIS Graduate student @ The University of Pennsylvania
