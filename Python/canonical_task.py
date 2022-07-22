#!/usr/bin/python3
import os
import gzip
from heapq import nlargest
import io
import sys


NUMBER_OF_RESULTS = 10

Package_Dict = {}

'''This program simply fetch file and downloads it to parse on later.
Parsing the results will provide us with the required/expected output'''

class DebianPacks():
    '''class'''
    def __init__(self):
        '''takes the argument passed from command line'''
        archi = sys.argv[1]
        # data to be saved in this fileformat
        file_name = "Contents-{arch}.gz".format(arch=archi)
        url = self.get_url(file_name)
        print("Downloading file: %s" % file_name)
        print("Downloading from: %s" % url)
        self.download_file(url)
        print("Parsing file: %s " % file_name)
        self.read_file(file_name)
    def download_file(self, file):
        '''using wget to fetch file'''
        wget_command = "wget {} --no-check-certificate".format(file)
        os.system(wget_command)
    def read_file(self, file):
        '''reading and parsing data'''
        gz = gzip.open(file, 'rb')
        f = io.BufferedReader(gz)
        for line in f:
            line = line.decode("utf-8")
            line = line.rstrip()
            file_name, space, package_name = line.rpartition(' ')
            package_name = package_name.split(',')
            for package in package_name:
                # get package name
                package = package.rpartition('/')[2]
                # Uniquessness in keys
                if package not in Package_Dict.keys():
                    Package_Dict[package] = []
                Package_Dict[package].append(file_name)
        gz.close()
        i = 0
        line = '-' * 100
        print(line)
        print('{:<20s}{:<30s}{:>20s}'.format("#", "Package", "Count"))
        print(line)
        #Heap queue is used for processing
        #heapq.nlargest provides results sorted order from largest to smallest
        for package in nlargest(NUMBER_OF_RESULTS,
                                Package_Dict, key=lambda e: len(Package_Dict[e])):
            print(
                "{:<20s}{: <30} {: >20}"
                .format("{}.".format(i+1), package, len(Package_Dict[package]))
            )
            i += 1            
    def get_url(self, filename):
        '''mirror'''
        url = "http://ftp.uk.debian.org/debian/dists/stable/main/"
        file_url = "{deb_mirror}{filename}".format(
            filename=filename,
            deb_mirror=url,
        )
        return file_url
if __name__ == "__main__":
    DebianPacks()
