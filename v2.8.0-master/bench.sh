#!/bin/bash
GO111MODULE=off go build
bin=v2.8.0-master
# SetSheetRow
./$bin -func=SetSheetRow -rows=200 -cols=50 -n=6
./$bin -func=SetSheetRow -rows=400 -cols=50 -n=6
./$bin -func=SetSheetRow -rows=800 -cols=50 -n=6
./$bin -func=SetSheetRow -rows=1600 -cols=50 -n=6
./$bin -func=SetSheetRow -rows=3200 -cols=50 -n=6
./$bin -func=SetSheetRow -rows=6400 -cols=50 -n=6
./$bin -func=SetSheetRow -rows=12800 -cols=50 -n=6
./$bin -func=SetSheetRow -rows=25600 -cols=50 -n=6
./$bin -func=SetSheetRow -rows=51200 -cols=50 -n=6
./$bin -func=SetSheetRow -rows=102400 -cols=50 -n=6
# StreamWriter
./$bin -func=StreamWriter -rows=200 -cols=50 -n=6
./$bin -func=StreamWriter -rows=400 -cols=50 -n=6
./$bin -func=StreamWriter -rows=800 -cols=50 -n=6
./$bin -func=StreamWriter -rows=1600 -cols=50 -n=6
./$bin -func=StreamWriter -rows=3200 -cols=50 -n=6
./$bin -func=StreamWriter -rows=6400 -cols=50 -n=6
./$bin -func=StreamWriter -rows=12800 -cols=50 -n=6
./$bin -func=StreamWriter -rows=25600 -cols=50 -n=6
./$bin -func=StreamWriter -rows=51200 -cols=50 -n=6
./$bin -func=StreamWriter -rows=102400 -cols=50 -n=6
# AddChart
./$bin -func=AddChart -rows=200 -cols=50
# SetCellHyperLink
./$bin -func=SetCellHyperLink -rows=200 -cols=50
./$bin -func=SetCellHyperLink -rows=400 -cols=50
./$bin -func=SetCellHyperLink -rows=800 -cols=50
# AddPicture
./$bin -func=AddPicture -rows=200 -cols=50
./$bin -func=AddPicture -rows=400 -cols=50
./$bin -func=AddPicture -rows=800 -cols=50
./$bin -func=AddPicture -rows=1600 -cols=50
# MergeCell
./$bin -func=MergeCell -n=200
./$bin -func=MergeCell -n=400
./$bin -func=MergeCell -n=800
./$bin -func=MergeCell -n=1600
./$bin -func=MergeCell -n=3200
./$bin -func=MergeCell -n=6400
# RowIterator
./$bin -func=RowIterator -rows=200 -cols=50
./$bin -func=RowIterator -rows=400 -cols=50
./$bin -func=RowIterator -rows=800 -cols=50
./$bin -func=RowIterator -rows=1600 -cols=50
./$bin -func=RowIterator -rows=3200 -cols=50
./$bin -func=RowIterator -rows=6400 -cols=50
./$bin -func=RowIterator -rows=12800 -cols=50
./$bin -func=RowIterator -rows=25600 -cols=50
./$bin -func=RowIterator -rows=51200 -cols=50
./$bin -func=RowIterator -rows=102400 -cols=50
# GetRows
./$bin -func=GetRows -rows=200 -cols=50
./$bin -func=GetRows -rows=400 -cols=50
./$bin -func=GetRows -rows=800 -cols=50
./$bin -func=GetRows -rows=1600 -cols=50
./$bin -func=GetRows -rows=3200 -cols=50
./$bin -func=GetRows -rows=6400 -cols=50
./$bin -func=GetRows -rows=12800 -cols=50
./$bin -func=GetRows -rows=25600 -cols=50
./$bin -func=GetRows -rows=51200 -cols=50
./$bin -func=GetRows -rows=102400 -cols=50
# Cleanup
rm -rf *.xlsx
