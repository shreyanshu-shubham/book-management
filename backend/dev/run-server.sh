#!/bin/sh
su - root -c "update-rc.d bk-manage defaults"
su - root -c "service bk-manage start"