#!/usr/bin/env python3
"""Run the existing ORM generator and normalize its legacy unkeyed test literals."""
import os, pathlib, re, subprocess, sys
root=pathlib.Path(__file__).resolve().parents[1]
env=os.environ.copy();env['RUN_TARANTOOL_GENERATOR']='1'
subprocess.run([sys.argv[1],'test','./mTest','-run','TestGenerateOrm'],cwd=root,env=env,check=True)
# Keep the generator's schema and benchmark operations intact. Only spell out
# external struct field names so regenerated tests pass current go vet checks.
for path in sorted((root/'mTest').rglob('*__ORM.GEN_test.go')):
 lines=path.read_text().splitlines(keepends=True);in_fields=False
 for i,line in enumerate(lines):
  if 'Fields: []Tt.Field{' in line:in_fields=True;continue
  if in_fields and line.strip()=='},':in_fields=False
  if in_fields:
   line=re.sub(r'^(\s*)\{([^{}]+?),\s*([^{}]+?)\},\s*$',r'\1{Name: \2, Type: \3},\n',line)
  lines[i]=re.sub(r'tarantool\.IntKey\{(\d+)\}',r'tarantool.IntKey{I: \1}',line)
 text=''.join(lines)
 if text!=path.read_text():path.write_text(text)
 subprocess.run(['gofmt','-w',str(path)],check=True)
