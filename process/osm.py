import osmium as osm
import numpy as np
import pandas as pd

# An obstacle contains nodes (lat, lon) which forms a polygon and has
# height >= 12

# OSMHandler extract restaurant's locations from the osm file and obstacles
class OSMHandler(osm.SimpleHandler):
    def __init__(self, region=None):
        osm.SimpleHandler.__init__(self)
        self.region = region
        self.osm_data = list()
        self.node_location_map = dict()

    # add nodes which are restaurants
    def node(self, n):
        location = OSMHandler.getLocation(n.location)
        self.node_location_map[n.id] = location

        if OSMHandler.is_restaurant(n):
            self.osm_data.append(location)

    # create csv file for shops and obstacles
    def save(self):
        # save restaurants coordinates
        data_columns = ['lat', 'lon']
        df = pd.DataFrame(self.osm_data, columns=data_columns)
        df.to_csv(f'csv/{self.region}/shops.csv', index=False)

    @staticmethod
    def is_restaurant(elem):
        for tag in elem.tags:
            if tag.v == 'restaurant':
                return True
        return False

    @staticmethod
    def getLocation(location):
        return np.array([float(location.lat), float(location.lon)])
