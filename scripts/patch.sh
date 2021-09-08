stitch() {
  echo finding pb files in "$1"
  pb_files=$(find "$1" -name "*.pb.go")
  for file in $pb_files; do
    echo "patching pb file $file"
    nana delete -func ProtoReflect -file "$file"
  done
}

for dir in "$@"
do
  echo stitching up "$dir"
  stitch "$dir"
done

echo "done stitching"