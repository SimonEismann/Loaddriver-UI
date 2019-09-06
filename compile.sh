cd load-director
docker build -t descartesresearch/web-load-driver .
cd ..
cd load-generator
docker build -t descartesresearch/web-load-generator .
cd ..