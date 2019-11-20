<?php

namespace App\Admin\Models;

use Illuminate\Database\Eloquent\Model;
use App\Observers\ModelObserver;


class Pool extends Model
{
    protected $table = 'pool';

    public function tags()
    {
        return $this->hasMany(PoolCoinbaseTag::class, "pool_id");
    }

    public function addresses()
    {
        return $this->hasMany(PoolAddress::class, "pool_id");
    }


    public static function boot()
    {
        parent::boot();
        $class = get_called_class();
        $class::observe(new ModelObserver());

        static::deleting(function ($pool) {
            /**
             * @var $pool Pool
             */
            $pool->tags()->delete();
            $pool->addresses()->delete();
        });
    }
}
